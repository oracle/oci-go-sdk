package common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultHostUrlTemplate   = "%s.%s.oraclecloud.com"
	defaultScheme            = "https"
	defaultSDKMarker         = "Oracle-GoSDK"
	defaultUserAgentTemplate = "%s/%s (%s/%s; go/%s)" //SDK/SDKVersion (OS/OSVersion; Lang/LangVersion)
	defaultTimeout           = time.Second * 30
)

type RequestInterceptor func(*http.Request) error

// HttpRequestor wraps the execution of a http request, it is generally implemented by
// http.Client.Do, but can be customized for testing
type HttpRequestDispatcher interface {
	Do(req *http.Request) (*http.Response, error)
}

//BaseClient struct implements all basic operations to call oci web services.
type BaseClient struct {
	//HttpClient performs the http network operations
	HttpClient HttpRequestDispatcher

	//Signer performs auth operation
	Signer HttpRequestSigner

	//A request interceptor can be used to customize the request before signing and dispatching
	Interceptor RequestInterceptor

	//The region of the client
	Region Region

	//The host of the service
	Host string

	//The user agent
	UserAgent string

	//Base path for all operations of this client
	BasePath string
}

func newClientWithHttpDispatcher(dispatcher HttpRequestDispatcher, region Region, provider ConfigurationProvider) (client BaseClient) {
	userAgent := fmt.Sprintf(defaultUserAgentTemplate, defaultSDKMarker, Version(), runtime.GOOS, runtime.GOARCH, runtime.Version())
	signer := NewOCIRequestSigner(provider)

	client = BaseClient{
		UserAgent:   userAgent,
		Region:      region,
		Interceptor: nil,
		Signer:      signer,
		HttpClient:  dispatcher,
	}
	return
}

func getDefaultHttpDispatcher() http.Client {
	httpClient := http.Client{
		Timeout:   defaultTimeout,
		Transport: &http.Transport{},
	}
	return httpClient
}

//Create a new client with a configuration provider, the configuration provider
//will be used for the default signer as well as reading the region
func NewClientWithConfig(configProvider ConfigurationProvider) (client BaseClient, err error) {
	regionstr, err := configProvider.Region()
	if err != nil {
		err = fmt.Errorf("can not create client, bad region configuration: %s", err.Error())
		return
	}
	region, err := StringToRegion(regionstr)
	if err != nil {
		err = fmt.Errorf("can not create client, bad region configuration: %s", err.Error())
		return
	}

	dispatcher := getDefaultHttpDispatcher()
	client = newClientWithHttpDispatcher(&dispatcher, region, configProvider)
	return
}

//Create a new default client for a given region
func NewClientForRegion(region Region) (client BaseClient) {
	dispatcher := getDefaultHttpDispatcher()
	provider := environmentConfigurationProvider{EnvironmentVariablePrefix: "TF_VAR"}
	return newClientWithHttpDispatcher(&dispatcher, region, provider)
}

func (client *BaseClient) prepareRequest(request *http.Request) (err error) {
	request.Header.Set("User-Agent", client.UserAgent)
	request.URL.Scheme = defaultScheme
	request.URL.Host = client.Host
	currentPath := request.URL.Path
	request.URL.Path = path.Clean(fmt.Sprintf("/%s/%s", client.BasePath, currentPath))
	return
}

func (client BaseClient) intercept(request *http.Request) (err error) {
	if client.Interceptor != nil {
		err = client.Interceptor(request)
	}
	return
}

func checkForSuccessfulResponse(res *http.Response) error {
	familyStatusCode := res.StatusCode / 100
	if familyStatusCode == 4 || familyStatusCode == 5 {
		return newServiceFailureFromResponse(res)
	}
	return nil

}

func (client BaseClient) Call(ctx context.Context, request *http.Request) (response *http.Response, err error) {
	Debugln("Atempting to call downstream service")
	request = request.WithContext(ctx)

	err = client.prepareRequest(request)
	if err != nil {
		return
	}

	//Intercept
	err = client.intercept(request)
	if err != nil {
		return
	}

	//Sign the request
	err = client.Signer.Sign(request)
	if err != nil {
		return
	}

	IfDebug(func() {
		if dump, e := httputil.DumpRequest(request, true); e == nil {
			Logf("Dump Request %v", string(dump))
		} else {
			Debugln(e)
		}
	})

	//Execute the http request
	response, err = client.HttpClient.Do(request)

	IfDebug(func() {
		if err != nil {
			Logln(err)
			return
		}

		if dump, e := httputil.DumpResponse(response, true); e == nil {
			Logf("Dump Response %v", string(dump))
		} else {
			Debugln(e)
		}
	})

	if err != nil {
		return
	}

	err = checkForSuccessfulResponse(response)
	return
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Request Marshaling
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func isNil(v reflect.Value) bool {
	return v.Kind() == reflect.Ptr && v.IsNil()
}

//Returns the string representation of a reflect.Value
//Only transforms primitive values
func toStringValue(v reflect.Value, field reflect.StructField) (string, error) {
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return "", fmt.Errorf("Can not marshal a nil pointer")
		}
		v = v.Elem()
	}

	if v.Type() == timeType {
		t := v.Interface().(SDKTime)
		return FormatTime(t), nil
	}

	switch v.Kind() {
	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.String:
		return v.String(), nil
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', 6, 32), nil
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', 6, 64), nil
	default:
		return "", fmt.Errorf("Marshaling structure to a http.Request does not support field named: %s of type: %v",
			field.Name, v.Type().String())
	}
}

func addToBody(request *http.Request, value reflect.Value, field reflect.StructField) (e error) {
	Debugln("Marshaling to body from field:", field.Name)
	if request.Body != nil {
		Logln("The body of the request is already set. Structure: ", field.Name, " will overwrite it")
	}
	marshaled, e := json.Marshal(value.Interface())
	if e != nil {
		return
	}
	Debugf("Marshaled body is: %s", string(marshaled))
	bodyBytes := bytes.NewReader(marshaled)
	request.ContentLength = int64(bodyBytes.Len())
	request.Header.Set("Content-Length", strconv.FormatInt(request.ContentLength, 10))
	request.Body = ioutil.NopCloser(bodyBytes)
	request.GetBody = func() (io.ReadCloser, error) {
		return ioutil.NopCloser(bodyBytes), nil
	}
	return
}

func addToQuery(request *http.Request, value reflect.Value, field reflect.StructField) (e error) {
	Debugln("Marshaling to query from field:", field.Name)
	if request.URL == nil {
		request.URL = &url.URL{}
	}
	query := request.URL.Query()
	var queryParameterValue, queryParameterName string

	if queryParameterName = field.Tag.Get("name"); queryParameterName == "" {
		return fmt.Errorf("Marshaling request to a query requires the 'name' tag for field: %s ", field.Name)
	}

	mandatory, _ := strconv.ParseBool(strings.ToLower(field.Tag.Get("mandatory")))

	//If mandatory and nil. Error out
	if mandatory && isNil(value) {
		return fmt.Errorf("marshaling request to a header requires not nil pointer for field: %s", field.Name)
	}

	//if not mandatory and nil. Omit
	if !mandatory && isNil(value) {
		Debugf("Query parameter value is not mandatory and is nil pointer in field: %s. Skipping header", field.Name)
		return
	}

	if queryParameterValue, e = toStringValue(value, field); e != nil {
		return
	}

	query.Set(queryParameterName, queryParameterValue)
	request.URL.RawQuery = query.Encode()
	return
}

//Adds to the path of the url in the order they appear in the structure
func addToPath(request *http.Request, value reflect.Value, field reflect.StructField) (e error) {
	var additionalUrlPathPart string
	if additionalUrlPathPart, e = toStringValue(value, field); e != nil {
		return fmt.Errorf("can not marshal to path in request for field %s. Due to %s", field.Name, e.Error())
	}

	if request.URL == nil {
		request.URL = &url.URL{}
		request.URL.Path = ""
	}
	var currentUrlPath = request.URL.Path

	var templatedPathRegex, _ = regexp.Compile(".*{.+}.*")
	if !templatedPathRegex.MatchString(currentUrlPath) {
		Debugln("Marshaling request to path by appending field:", field.Name)
		allPath := []string{currentUrlPath, additionalUrlPathPart}
		newPath := strings.Join(allPath, "/")
		request.URL.Path = path.Clean(newPath)
		return
	} else {
		var fieldName string
		if fieldName = field.Tag.Get("name"); fieldName == "" {
			e = fmt.Errorf("Marshaling request to path name and template requires a 'name' tag for field: %s", field.Name)
			return
		}
		urlTemplate := currentUrlPath
		Debugln("Marshaling to path from field:", field.Name, "in template:", urlTemplate)
		request.URL.Path = path.Clean(strings.Replace(urlTemplate, "{"+fieldName+"}", additionalUrlPathPart, -1))
		return
	}
}

func addToHeader(request *http.Request, value reflect.Value, field reflect.StructField) (e error) {
	Debugln("Marshaling to header from field:", field.Name)
	if request.Header == nil {
		request.Header = http.Header{}
	}

	var headerName, headerValue string
	if headerName = field.Tag.Get("name"); headerName == "" {
		return fmt.Errorf("marshaling request to a header requires the 'name' tag for field: %s", field.Name)
	}

	mandatory, _ := strconv.ParseBool(strings.ToLower(field.Tag.Get("mandatory")))
	//If mandatory and nil. Error out
	if mandatory && isNil(value) {
		return fmt.Errorf("marshaling request to a header requires not nil pointer for field: %s", field.Name)
	}

	//if not mandatory and nil. Omit
	if !mandatory && isNil(value) {
		Debugf("Header value is not mandatory and is nil pointer in field: %s. Skipping header", field.Name)
		return
	}

	//Otherwise get value and set header
	if headerValue, e = toStringValue(value, field); e != nil {
		return
	}

	request.Header.Set(headerName, headerValue)
	return
}

//Makes sure the incoming structure is able to be marshalled
//to a request
func checkForValidRequestStruct(s interface{}) (*reflect.Value, error) {
	val := reflect.ValueOf(s)
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil, fmt.Errorf("Can not marshal to request a pointer to structure")
		}
		val = val.Elem()
	}

	if s == nil {
		return nil, fmt.Errorf("Can not marshal to request a nil structure")
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Can not marshal to request, expects struct input. Got %v", val.Kind())
	}

	return &val, nil
}

// Populates the parts of a request by reading tags in the passed structure
// nested structs are followed recursively depth-first.
func structToRequestPart(request *http.Request, val reflect.Value) (err error) {
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		if err != nil {
			return
		}

		sf := typ.Field(i)
		//unexported
		if sf.PkgPath != "" && !sf.Anonymous {
			continue
		}

		sv := val.Field(i)
		tag := sf.Tag.Get("contributesTo")
		switch tag {
		case "header":
			err = addToHeader(request, sv, sf)
		case "path":
			err = addToPath(request, sv, sf)
		case "query":
			err = addToQuery(request, sv, sf)
		case "body":
			err = addToBody(request, sv, sf)
		case "":
			Debugln(sf.Name, "does not contain contributes tag. Skipping.")
		default:
			err = fmt.Errorf("Can not marshal field: %s. It needs to contain valid contributesTo tag", sf.Name)
		}
	}
	return
}

// Marshals a structure to an http request using tag values in the struct
// The marshaller tag should like the following
// type A struct {
// 		 ANumber string `contributesTo="query" name="number"`
//		 TheBody `contributesTo="body"`
// }
// where the contributesTo tag can be: header, path, query, body
// and the 'name' tag is the name of the value used in the http request(not applicable for path)
// If path is specified as part of the tag, the values are appened to the url path
// in the order they appear in the structure
// The current implementation only supports primitive types, except for the body tag, which needs a struct type.
// The body of a request will be marshaled using the tags of the structure
func HttpRequestMarshaller(requestStruct interface{}, httpRequest *http.Request) (err error) {
	var val *reflect.Value
	if val, err = checkForValidRequestStruct(requestStruct); err != nil {
		return
	}

	Debugln("Marshaling to Request:", val.Type().Name())
	err = structToRequestPart(httpRequest, *val)
	return
}

// MakeDefaultHttpRequest creates the basic http request with the necessary headers set
func MakeDefaultHttpRequest(method, path string) (httpRequest http.Request) {
	httpRequest = http.Request{
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		URL:        &url.URL{},
	}

	httpRequest.Header.Set("Content-Length", "0")
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Date", time.Now().UTC().Format(http.TimeFormat))
	httpRequest.Header.Set("Opc-Client-Info", strings.Join([]string{defaultSDKMarker, Version()}, "/"))
	httpRequest.Header.Set("Accept", "*/*")
	httpRequest.Method = method
	httpRequest.URL.Path = path
	return
}

// MakeDefaultHttpRequestWithTaggedStruct creates an http request from an struct with tagged fields, see HttpRequestMarshaller
// for more information
func MakeDefaultHttpRequestWithTaggedStruct(method, path string, requestStruct interface{}) (httpRequest http.Request, err error) {
	httpRequest = MakeDefaultHttpRequest(method, path)
	err = HttpRequestMarshaller(requestStruct, &httpRequest)
	return
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Request UnMarshaling
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//Makes sure the incoming structure is able to be unmarshaled
//to a request
func checkForValidResponseStruct(s interface{}) (*reflect.Value, error) {
	val := reflect.ValueOf(s)
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil, fmt.Errorf("can not unmarshal to response a pointer to nil structure")
		}
		val = val.Elem()
	}

	if s == nil {
		return nil, fmt.Errorf("can not unmarshal to response a nil structure")
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("can not unmarshal to response, expects struct input. Got %v", val.Kind())
	}

	return &val, nil
}

func intSizeFromKind(kind reflect.Kind) int {
	switch kind {
	case reflect.Int8, reflect.Uint8:
		return 8
	case reflect.Int16, reflect.Uint16:
		return 16
	case reflect.Int32, reflect.Uint32:
		return 32
	case reflect.Int64, reflect.Uint64:
		return 64
	case reflect.Int, reflect.Uint:
		return strconv.IntSize
	default:
		Debugln("The type is not valid: %v. Returing int size for arch", kind.String())
		return strconv.IntSize
	}

}

func analyzeValue(stringValue string, kind reflect.Kind) (val reflect.Value, valPointer reflect.Value, err error) {
	switch kind {
	case timeType.Kind():
		var t time.Time
		t, err = time.Parse(time.RFC3339, stringValue)
		if err != nil {
			return
		}
		sdkTime := SDKTimeFromTime(t)
		val = reflect.ValueOf(sdkTime)
		valPointer = reflect.ValueOf(&sdkTime)
		return
	case reflect.Bool:
		var bVal bool
		if bVal, err = strconv.ParseBool(stringValue); err != nil {
			return
		}
		val = reflect.ValueOf(bVal)
		valPointer = reflect.ValueOf(&bVal)
		return
	case reflect.Int:
		size := intSizeFromKind(kind)
		var iVal int64
		if iVal, err = strconv.ParseInt(stringValue, 10, size); err != nil {
			return
		}
		var iiVal int
		iiVal = int(iVal)
		val = reflect.ValueOf(iiVal)
		valPointer = reflect.ValueOf(&iiVal)
		return
	case reflect.Int64:
		size := intSizeFromKind(kind)
		var iVal int64
		if iVal, err = strconv.ParseInt(stringValue, 10, size); err != nil {
			return
		}
		val = reflect.ValueOf(iVal)
		valPointer = reflect.ValueOf(&iVal)
		return
	case reflect.Uint:
		size := intSizeFromKind(kind)
		var iVal uint64
		if iVal, err = strconv.ParseUint(stringValue, 10, size); err != nil {
			return
		}
		var uiVal uint
		uiVal = uint(iVal)
		val = reflect.ValueOf(uiVal)
		valPointer = reflect.ValueOf(&uiVal)
		return
	case reflect.String:
		val = reflect.ValueOf(stringValue)
		valPointer = reflect.ValueOf(&stringValue)
	case reflect.Float32:
		var fVal float64
		if fVal, err = strconv.ParseFloat(stringValue, 32); err != nil {
			return
		}
		var ffVal float32
		ffVal = float32(fVal)
		val = reflect.ValueOf(ffVal)
		valPointer = reflect.ValueOf(&ffVal)
		return
	case reflect.Float64:
		var fVal float64
		if fVal, err = strconv.ParseFloat(stringValue, 64); err != nil {
			return
		}
		val = reflect.ValueOf(fVal)
		valPointer = reflect.ValueOf(&fVal)
		return
	default:
		err = fmt.Errorf("Value for kind: %s not supported", kind)
	}
	return
}

//Sets the field of a struct, with the appropiate value of the string
//Only sets basic types
func fromStringValue(newValue string, val *reflect.Value, field reflect.StructField) (err error) {

	if !val.CanSet() {
		err = fmt.Errorf("can not set field name: %s of type: %v", field.Name, val.Type().String())
		return
	}

	kind := val.Kind()
	isPointer := false
	if val.Kind() == reflect.Ptr {
		isPointer = true
		kind = field.Type.Elem().Kind()
	}

	value, valPtr, err := analyzeValue(newValue, kind)
	if err != nil {
		return
	}
	if !isPointer {
		val.Set(value)
	} else {
		val.Set(valPtr)
	}
	return
}

func addFromBody(response *http.Response, value *reflect.Value, field reflect.StructField) (err error) {
	Debugln("Unmarshaling from body to field:", field.Name)
	if response.Body == nil {
		Debugln("Unmarshaling body skipped due to nil body content for field: ", field.Name)
		return nil
	}

	//Consumes the body, consider implementing it
	//without body consumption
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	newStruct := reflect.New(value.Type()).Interface()
	err = json.Unmarshal(content, &newStruct)
	if err != nil {
		return
	}

	newVal := reflect.ValueOf(newStruct)
	if newVal.Kind() == reflect.Ptr {
		newVal = newVal.Elem()
	}
	value.Set(newVal)
	return

}

func addFromHeader(response *http.Response, value *reflect.Value, field reflect.StructField) (err error) {
	Debugln("Unmarshaling from header to field:", field.Name)
	var headerName string
	if headerName = field.Tag.Get("name"); headerName == "" {
		return fmt.Errorf("Unmarshaling response to a header requires the 'name' tag for field: %s", field.Name)
	}

	headerValue := response.Header.Get(headerName)
	if headerValue == "" {
		Debugf("Unmarshalling did not find header with name:%s", headerName)
		return nil
	}

	if err = fromStringValue(headerValue, value, field); err != nil {
		return fmt.Errorf("unmarshaling response to a header failed for field %s, due to %s", field.Name,
			err.Error())
	}
	return
}

// Populates a struct from parts of a request by reading tags of the struct
func responseToStruct(response *http.Response, val *reflect.Value) (err error) {
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		if err != nil {
			return
		}

		sf := typ.Field(i)

		//unexported
		if sf.PkgPath != "" {
			continue
		}

		sv := val.Field(i)
		tag := sf.Tag.Get("presentIn")
		switch tag {
		case "header":
			err = addFromHeader(response, &sv, sf)
		case "body":
			err = addFromBody(response, &sv, sf)
		case "":
			Debugln(sf.Name, "does not contain presentIn tag. Skipping")
		default:
			err = fmt.Errorf("can not unmarshal field: %s. It needs to contain valid presentIn tag", sf.Name)
		}
	}
	return
}

// UnmarshalResponse hydrates the fileds of an struct with the values of an http response, guided
// by the field tags. The directive tag is "presentIn" and it can be either
//  - "header": Will look for the header tagged as "name" in the headers of the struct and set it value to that
//  - "body": It will try to marshal the body from a json string to a struct tagged with 'presentIn: "body"'.
// Further this method will consume the body as it should safe to close after it returns
// Notice the current implementation only supports native types:int, strings, floats, bool
func UnmarshalResponse(httpResponse *http.Response, responseStruct interface{}) (err error) {
	var val *reflect.Value
	if val, err = checkForValidResponseStruct(responseStruct); err != nil {
		return
	}

	if err = responseToStruct(httpResponse, val); err != nil {
		return
	}

	return nil
}
