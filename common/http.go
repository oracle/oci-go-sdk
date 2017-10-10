package common

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	"io"
)

const (
	defaultHostUrlTemplate   = "%s.%s.oraclecloud.com"
	defaultScheme            = "https"
	defaultSDKMarker         = "Oracle-GoSDK"
	defaultUserAgentTemplate = "%s/%s (%s/%s; go/%s)" //SDK/SDKVersion (OS/OSVersion; Lang/LangVersion)
	defaultTimeout           = time.Second * 15
)

type RequestInterceptor func(*http.Request) error

// HttpRequestor wraps the execution of a http request, it is generally implemented by
// http.Client.Do, but can be customized for testing
type HttpRequestDispatcher interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client wraps the basic operation to call the downstream services
type Client interface {
	Call(request http.Request) (response *http.Response, err error)
}

//BaseClient struct implements all basic operations to call oci web services.
type BaseClient struct {
	httpClient            HttpRequestDispatcher
	Signer                RequestSigner
	ApiVersion            string
	UserAgent             string
	ServiceName           string
	Region                Region
	ConfigurationProvider ConfigurationProvider
	//A request interceptor can be used to customize the request before signing and dispatching
	Interceptor RequestInterceptor
}

func NewClientWithHttpDispatcher(dispatcher HttpRequestDispatcher) (client BaseClient) {
	userAgent := fmt.Sprintf(defaultUserAgentTemplate, defaultSDKMarker, Version(), runtime.GOOS, runtime.GOARCH, runtime.Version())
	provider := EnvironmentConfigurationProvider{EnvironmentVariablePrefix: "TF_VAR"}

	client = BaseClient{
		UserAgent:             userAgent,
		Region:                DefaultRegion,
		Interceptor:           nil,
		ConfigurationProvider: provider,
		Signer:                OCIRequestSigner{KeyProvider: provider},
		httpClient:            dispatcher,
	}
	return
}

func NewClient() (client BaseClient) {
	return NewClientWithHttpDispatcher(&http.Client{
		Timeout:   defaultTimeout,
		Transport: &http.Transport{},
	})
}

func (client *BaseClient) prepareRequest(request *http.Request) (err error) {
	regionString, err := RegionToString(client.Region)
	if err != nil {
		return
	}
	request.Header.Set("User-Agent", client.UserAgent)
	hostUrl := fmt.Sprintf(defaultHostUrlTemplate, client.ServiceName, regionString)
	request.URL.Host = hostUrl
	request.URL.Scheme = defaultScheme
	currentPath := request.URL.Path
	request.URL.Path = path.Clean(fmt.Sprintf("%s/%s", client.ApiVersion, currentPath))
	return
}

func (client BaseClient) intercept(request *http.Request) (err error) {
	if client.Interceptor != nil {
		err = client.Interceptor(request)
	}
	return
}

func calculateHashOfBody(request *http.Request) (err error) {
	if request.Method == http.MethodPost || request.Method == http.MethodPut {
		if request.ContentLength > 0 {
			hash, e := GetBodyHash(*request)
			if e != nil {
				return e
			}
			request.Header.Set("X-Content-Sha256", hash)
		}
	}
	return
}

func (client BaseClient) Call(request http.Request) (response *http.Response, err error) {
	Debugln("Atempting to call downstream service")
	err = client.prepareRequest(&request)
	if err != nil {
		return
	}

	//Intercept
	err = client.intercept(&request)
	if err != nil {
		return
	}


	//Sign the request
	err = client.Signer.Sign(&request)
	if err != nil {
		return
	}

	IfDebug(func() {
		if dump, e := httputil.DumpRequest(&request, true); e == nil {
			Logf("Dump Request %v", string(dump))
		} else {
			Debugln(e)
		}
	})

	response, err = client.httpClient.Do(&request)
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
	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Request Marshaling
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var timeType = reflect.TypeOf(time.Time{})

func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
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
		t := v.Interface().(time.Time)
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
		return
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
		return fmt.Errorf("Marshaling request to a header requires the 'name' tag for field: %s", field.Name)
	}
	if headerValue, e = toStringValue(value, field); e != nil {
		return
	}
	header := request.Header
	header.Set(headerName, headerValue)
	return
}

//Makes sure the incoming structure is able to be marshalled
//to a request
func checkForValidStruct(s interface{}) (*reflect.Value, error) {
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
	if val, err = checkForValidStruct(requestStruct); err != nil {
		return
	}

	Debugln("Marshaling to Request:", val.Type().Name())
	err = structToRequestPart(httpRequest, *val)
	return
}

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

func MakeDefaultHttpRequestWithTaggedStruct(method, path string, requestStruct interface{}) (httpRequest http.Request, err error) {
	httpRequest = MakeDefaultHttpRequest(method, path)
	err = HttpRequestMarshaller(requestStruct, &httpRequest)
	return
}
