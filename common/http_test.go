package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Test data structures, avoid import cycle
type TestupdateUserDetails struct {
	Description string  `mandatory:"false" json:"description,omitempty"`
	Name        *string `mandatory:"false" json:"name"`
	SomeNumbers []int   `mandatory:"false" json:"numbers"`
}

type listCompartmentsRequest struct {
	CompartmentID string `mandatory:"true" contributesTo:"query" name:"compartmentId"`
	Page          string `mandatory:"false" contributesTo:"query" name:"page"`
	Limit         int32  `mandatory:"false" contributesTo:"query" name:"limit"`
}

type updateUserRequest struct {
	UserID                string `mandatory:"true" contributesTo:"path" name:"userId"`
	TestupdateUserDetails `contributesTo:"body"`
	IfMatch               string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

type TestcreateAPIKeyDetails struct {
	Key string `mandatory:"true" json:"key,omitempty"`
}

type TestcreateAPIKeyDetailsPtr struct {
	Key     *string  `mandatory:"true" json:"key,omitempty"`
	TheTime *SDKTime `mandatory:"true" json:"theTime,omitempty"`
}

type uploadAPIKeyRequest struct {
	UserID                  string `mandatory:"true" contributesTo:"path" name:"userId"`
	TestcreateAPIKeyDetails `contributesTo:"body"`
	OpcRetryToken           string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`
}

type uploadAPIKeyRequestPtr struct {
	UserID                     *string `mandatory:"true" contributesTo:"path" name:"userId"`
	TestcreateAPIKeyDetailsPtr `contributesTo:"body"`
	OpcRetryToken              *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestHttpMarshallerInvalidStruct(t *testing.T) {
	request := http.Request{}
	err := HTTPRequestMarshaller("asdf", &request)
	assert.Error(t, err, nil)
}

func TestHttpRequestMarshallerQuery(t *testing.T) {
	s := listCompartmentsRequest{CompartmentID: "ocid1", Page: "p", Limit: 23}
	request := MakeDefaultHTTPRequest(http.MethodPost, "/")
	HTTPRequestMarshaller(s, &request)
	query := request.URL.Query()
	assert.True(t, query.Get("compartmentId") == "ocid1")
	assert.True(t, query.Get("page") == "p")
	assert.True(t, query.Get("limit") == "23")
}

func TestMakeDefault(t *testing.T) {
	r := MakeDefaultHTTPRequest(http.MethodPost, "/one/two")
	assert.NotEmpty(t, r.Header.Get("Date"))
	assert.NotEmpty(t, r.Header.Get("Opc-Client-Info"))
}

func TestHttpMarshallerSimpleHeader(t *testing.T) {
	s := updateUserRequest{UserID: "id1", IfMatch: "n=as", TestupdateUserDetails: TestupdateUserDetails{Description: "name of"}}
	request := MakeDefaultHTTPRequest(http.MethodPost, "/random")
	HTTPRequestMarshaller(s, &request)
	header := request.Header
	assert.True(t, header.Get("if-match") == "n=as")
}

func TestHttpMarshallerSimpleStruct(t *testing.T) {
	s := uploadAPIKeyRequest{UserID: "111", OpcRetryToken: "token", TestcreateAPIKeyDetails: TestcreateAPIKeyDetails{Key: "thekey"}}
	request := MakeDefaultHTTPRequest(http.MethodPost, "/random")
	HTTPRequestMarshaller(s, &request)
	assert.True(t, strings.Contains(request.URL.Path, "111"))
}

func TestHttpMarshallerSimpleBody(t *testing.T) {
	desc := "theDescription"
	s := updateUserRequest{UserID: "id1", IfMatch: "n=as", TestupdateUserDetails: TestupdateUserDetails{Description: desc}}
	request := MakeDefaultHTTPRequest(http.MethodPost, "/random")
	HTTPRequestMarshaller(s, &request)
	body, _ := ioutil.ReadAll(request.Body)
	var content map[string]string
	json.Unmarshal(body, &content)
	assert.Contains(t, content, "description")
	assert.Contains(t, content, "numbers")
	assert.Contains(t, content, "name")
	assert.Equal(t, "", content["name"])
	assert.Equal(t, "", content["numbers"])

	if val, ok := content["description"]; !ok || val != desc {
		assert.Fail(t, "Should contain: "+desc)
	}

}

func TestHttpMarshalerAll(t *testing.T) {
	desc := "theDescription"
	s := struct {
		ID      string                `contributesTo:"path"`
		Name    string                `contributesTo:"query" name:"name"`
		When    *SDKTime              `contributesTo:"query" name:"when"`
		Income  float32               `contributesTo:"query" name:"income"`
		Male    bool                  `contributesTo:"header" name:"male"`
		Details TestupdateUserDetails `contributesTo:"body"`
	}{
		"101", "tapir", now(), 3.23, true, TestupdateUserDetails{Description: desc},
	}
	request := MakeDefaultHTTPRequest(http.MethodPost, "/")
	e := HTTPRequestMarshaller(s, &request)
	assert.NoError(t, e)
	var content map[string]string
	body, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(body, &content)
	when := s.When.Format(time.RFC3339)
	assert.True(t, request.URL.Path == "/101")
	assert.True(t, request.URL.Query().Get("name") == s.Name)
	assert.True(t, request.URL.Query().Get("income") == strconv.FormatFloat(float64(s.Income), 'f', 6, 32))
	assert.True(t, request.URL.Query().Get("when") == when)
	assert.Contains(t, content, "description")
	assert.Equal(t, request.Header.Get("Content-Type"), "application/json")
	if val, ok := content["description"]; !ok || val != desc {
		assert.Fail(t, "Should contain: "+desc)
	}
}

func TestHttpMarshalerPointers(t *testing.T) {

	n := new(string)
	*n = "theName"
	s := struct {
		Name *string `contributesTo:"query" name:"name"`
	}{
		n,
	}
	request := MakeDefaultHTTPRequest(http.MethodPost, "/random")
	HTTPRequestMarshaller(s, &request)
	assert.NotNil(t, request)
	assert.True(t, request.URL.Query().Get("name") == *s.Name)
}

func TestHttpMarshalerPointersErrorHeader(t *testing.T) {

	n := new(string)
	*n = "theName"
	s := struct {
		Name *string `mandatory:"true" contributesTo:"header" name:"name"`
	}{
		nil,
	}
	request := MakeDefaultHTTPRequest(http.MethodPost, "/random")
	e := HTTPRequestMarshaller(s, &request)
	assert.Error(t, e)
}

func TestHttpMarshalerPointersErrorPath(t *testing.T) {

	n := new(string)
	*n = "theName"
	s := struct {
		Name *string `mandatory:"true" contributesTo:"path" name:"name"`
	}{
		nil,
	}
	request := MakeDefaultHTTPRequest(http.MethodPost, "/random")
	e := HTTPRequestMarshaller(s, &request)
	assert.Error(t, e)
}

func TestHttpMarshallerSimpleStructPointers(t *testing.T) {
	now := SDKTime{time.Now()}
	s := uploadAPIKeyRequestPtr{
		UserID:        String("111"),
		OpcRetryToken: nil,
		TestcreateAPIKeyDetailsPtr: TestcreateAPIKeyDetailsPtr{
			Key:     String("thekey"),
			TheTime: &now,
		}}
	request := MakeDefaultHTTPRequest(http.MethodPost, "/random")
	HTTPRequestMarshaller(s, &request)
	all, _ := ioutil.ReadAll(request.Body)
	assert.True(t, len(all) > 2)
	assert.Equal(t, "", request.Header.Get("opc-retry-token"))
	assert.True(t, strings.Contains(request.URL.Path, "111"))
	assert.True(t, strings.Contains(string(all), "thekey"))
	assert.Contains(t, string(all), now.Format(time.RFC3339))
}

func TestHttpMarshallerSimpleStructPointersFilled(t *testing.T) {
	s := uploadAPIKeyRequestPtr{
		UserID:                     String("111"),
		OpcRetryToken:              String("token"),
		TestcreateAPIKeyDetailsPtr: TestcreateAPIKeyDetailsPtr{Key: String("thekey")}}
	request := MakeDefaultHTTPRequest(http.MethodPost, "/random")
	HTTPRequestMarshaller(s, &request)
	assert.Equal(t, "token", request.Header.Get("opc-retry-token"))
	assert.True(t, strings.Contains(request.URL.Path, "111"))

}

func TestHttpMarshalerUntaggedFields(t *testing.T) {
	s := struct {
		Name  string `contributesTo:"query" name:"name"`
		AList []string
		AMap  map[string]int
		TestupdateUserDetails
	}{
		"theName", []string{"a", "b"}, map[string]int{"a": 1, "b": 2},
		TestupdateUserDetails{Description: "n"},
	}
	request := &http.Request{}
	e := HTTPRequestMarshaller(s, request)
	assert.NoError(t, e)
	assert.NotNil(t, request)
	assert.True(t, request.URL.Query().Get("name") == s.Name)
}
func TestHttpMarshalerPathTemplate(t *testing.T) {
	urlTemplate := "/name/{userId}/aaa"
	s := uploadAPIKeyRequest{UserID: "111", OpcRetryToken: "token", TestcreateAPIKeyDetails: TestcreateAPIKeyDetails{Key: "thekey"}}
	request := MakeDefaultHTTPRequest(http.MethodPost, urlTemplate)
	e := HTTPRequestMarshaller(s, &request)
	assert.NoError(t, e)
	assert.Equal(t, "/name/111/aaa", request.URL.Path)
}

func TestHttpMarshalerFunnyTags(t *testing.T) {
	s := struct {
		Name  string `contributesTo:"quer" name:"name"`
		AList []string
		AMap  map[string]int
		TestupdateUserDetails
	}{
		"theName", []string{"a", "b"}, map[string]int{"a": 1, "b": 2},
		TestupdateUserDetails{Description: "n"},
	}
	request := &http.Request{}
	e := HTTPRequestMarshaller(s, request)
	assert.Error(t, e)
}

func TestHttpMarshalerUnsupportedTypes(t *testing.T) {
	s1 := struct {
		Name string         `contributesTo:"query" name:"name"`
		AMap map[string]int `contributesTo:"query" name:"theMap"`
	}{
		"theName", map[string]int{"a": 1, "b": 2},
	}
	s2 := struct {
		Name  string   `contributesTo:"query" name:"name"`
		AList []string `contributesTo:"query" name:"theList"`
	}{
		"theName", []string{"a", "b"},
	}
	s3 := struct {
		Name                  string `contributesTo:"query" name:"name"`
		TestupdateUserDetails `contributesTo:"query" name:"str"`
	}{
		"theName", TestupdateUserDetails{Description: "a"},
	}
	n := new(string)
	col := make([]int, 10)
	*n = "theName"
	s4 := struct {
		Name *string `contributesTo:"query" name:"name"`
		Coll *[]int  `contributesTo:"query" name:"coll"`
	}{
		n, &col,
	}

	lst := []interface{}{s1, s2, s3, s4}
	for _, l := range lst {
		request := &http.Request{}
		e := HTTPRequestMarshaller(l, request)
		Debugln(e)
		assert.Error(t, e)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Response Unmarshaling
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ListRegionsResponse wrapper for the ListRegions operation
type listRegionsResponse struct {

	// The []Region instance
	Items []int `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string `presentIn:"header" name:"opcrequestid"`
}

type listUsersResponse struct {
	Items        []int   `presentIn:"body"`
	OpcRequestID string  `presentIn:"header" name:"opcrequestid"`
	OpcNextPage  int     `presentIn:"header" name:"opcnextpage"`
	SomeUint     uint    `presentIn:"header" name:"someuint"`
	SomeBool     bool    `presentIn:"header" name:"somebool"`
	SomeTime     SDKTime `presentIn:"header" name:"sometime"`
	SomeFloat    float64 `presentIn:"header" name:"somefloat"`
}

func TestUnmarshalResponse_StringHeader(t *testing.T) {
	header := http.Header{}
	opcID := "111"
	header.Set("OpcrequestId", opcID)
	r := http.Response{Header: header}
	s := listRegionsResponse{}
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.Equal(t, s.OpcRequestID, opcID)

}

func TestUnmarshalResponse_MixHeader(t *testing.T) {
	header := http.Header{}
	opcID := "111"
	nextPage := int(333)
	someuint := uint(12)
	somebool := true
	sometime := now()
	somefloat := 2.556

	header.Set("OpcrequestId", opcID)
	header.Set("opcnextpage", strconv.FormatInt(int64(nextPage), 10))
	header.Set("someuint", strconv.FormatUint(uint64(someuint), 10))
	header.Set("somebool", strconv.FormatBool(somebool))
	header.Set("sometime", formatTime(*sometime))
	header.Set("somefloat", strconv.FormatFloat(somefloat, 'f', 3, 64))

	r := http.Response{Header: header}
	s := listUsersResponse{}
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.Equal(t, s.OpcRequestID, opcID)
	assert.Equal(t, nextPage, s.OpcNextPage)
	assert.Equal(t, someuint, s.SomeUint)
	assert.Equal(t, somebool, s.SomeBool)
	assert.Equal(t, sometime.Format(time.RFC3339), s.SomeTime.Format(time.RFC3339))

}

type rgn struct {
	Key  string `mandatory:"false" json:"key,omitempty"`
	Name string `mandatory:"false" json:"name,omitempty"`
}

func TestUnmarshalResponse_SimpleBody(t *testing.T) {
	sampleResponse := `{"key" : "RegionFRA","name" : "eu-frankfurt-1"}`
	header := http.Header{}
	opcID := "111"
	header.Set("OpcrequestId", opcID)
	s := struct {
		Rg rgn `presentIn:"body"`
	}{}
	r := http.Response{Header: header}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	r.Body = ioutil.NopCloser(bodyBuffer)
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.Equal(t, "eu-frankfurt-1", s.Rg.Name)
}

func TestUnmarshalResponse_SimpleBodyList(t *testing.T) {
	sampleResponse := `[{"key" : "RegionFRA","name" : "eu-frankfurt-1"},{"key" : "RegionIAD","name" : "us-ashburn-1"}]`
	header := http.Header{}
	opcID := "111"
	header.Set("OpcrequestId", opcID)
	s := struct {
		Items []rgn `presentIn:"body"`
	}{}
	r := http.Response{Header: header}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	r.Body = ioutil.NopCloser(bodyBuffer)
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.NotEmpty(t, s.Items)
	assert.Equal(t, "eu-frankfurt-1", s.Items[0].Name)
	assert.Equal(t, "RegionIAD", s.Items[1].Key)
}

func TestUnmarshalResponse_SimpleBodyPtr(t *testing.T) {
	sampleResponse := `{"key" : "RegionFRA","name" : "eu-frankfurt-1"}`
	header := http.Header{}
	opcID := "111"
	header.Set("OpcrequestId", opcID)
	s := struct {
		Rg *rgn `presentIn:"body"`
	}{}
	r := http.Response{Header: header}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	r.Body = ioutil.NopCloser(bodyBuffer)
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.Equal(t, "eu-frankfurt-1", s.Rg.Name)
}

type testRnUnexported struct {
	Key  string `mandatory:"false" json:"key,omitempty"`
	Name string `mandatory:"false" json:"name,omitempty"`
}

type TestRn struct {
	Key  string `mandatory:"false" json:"key,omitempty"`
	Name string `mandatory:"false" json:"name,omitempty"`
}

type listRgRes struct {
	testRnUnexported `presentIn:"body"`
	OpcRequestID     string `presentIn:"header" name:"opcrequestid"`
}

type listRgResEx struct {
	TestRn       `presentIn:"body"`
	OpcRequestID string `presentIn:"header" name:"opcrequestid"`
}

type listRgResPtr struct {
	OpcRequestID               *string  `presentIn:"header" name:"opcrequestid"`
	NumericHeader              *int     `presentIn:"header" name:"numeric"`
	SomeTime                   *SDKTime `presentIn:"header" name:"theTime"`
	SomeBool                   *bool    `presentIn:"header" name:"aBool"`
	SomeUint                   *uint    `presentIn:"header" name:"aUint"`
	SomeFloat                  *float32 `presentIn:"header" name:"aFloat"`
	TestcreateAPIKeyDetailsPtr `presentIn:"body"`
}

func TestUnmarshalResponse_BodyAndHeaderUnex(t *testing.T) {
	sampleResponse := `{"key" : "RegionFRA","name" : "eu-frankfurt-1"}`
	header := http.Header{}
	opcID := "111"
	header.Set("OpcrequestId", opcID)
	s := listRgRes{}
	r := http.Response{Header: header}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	r.Body = ioutil.NopCloser(bodyBuffer)
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.Equal(t, opcID, s.OpcRequestID)
	assert.Equal(t, "", s.Name)
	assert.Equal(t, "", s.Key)
}

func TestUnmarshalResponse_BodyAndHeader(t *testing.T) {
	sampleResponse := `{"key" : "RegionFRA","name" : "eu-frankfurt-1"}`
	header := http.Header{}
	opcID := "111"
	header.Set("OpcrequestId", opcID)
	s := listRgResEx{}
	r := http.Response{Header: header}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	r.Body = ioutil.NopCloser(bodyBuffer)
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.Equal(t, opcID, s.OpcRequestID)
	assert.Equal(t, "eu-frankfurt-1", s.Name)
	assert.Equal(t, "RegionFRA", s.Key)
}

func TestUnmarshalResponse_BodyAndHeaderPtr(t *testing.T) {
	header := http.Header{}
	opcID := "111"
	numericHeader := "1414"
	someFloat := float32(2.332342)
	someUint := uint(33)
	theTime := SDKTime{time.Now()}
	theTimeStr := theTime.Format(sdkTimeFormat)
	sampleResponse := fmt.Sprintf(`{"key" : "RegionFRA","theTime" : "%s"}`, theTimeStr)
	header.Set("OpcrequestId", opcID)
	header.Set("numeric", numericHeader)
	header.Set("theTime", theTimeStr)
	header.Set("aBool", "true")
	header.Set("aUint", "33")
	header.Set("aFloat", "2.332342")
	s := listRgResPtr{}
	r := http.Response{Header: header}
	bodyBuffer := bytes.NewBufferString(sampleResponse)
	r.Body = ioutil.NopCloser(bodyBuffer)
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.Equal(t, opcID, *s.OpcRequestID)
	delta, _ := time.ParseDuration("1s")
	assert.WithinDuration(t, theTime.Time, s.SomeTime.Time, delta)
	assert.Equal(t, true, *s.SomeBool)
	assert.Equal(t, someFloat, *s.SomeFloat)
	assert.Equal(t, someUint, *s.SomeUint)
	assert.WithinDuration(t, theTime.Time, s.TheTime.Time, delta)
	assert.Equal(t, "RegionFRA", *s.Key)
}

type reqWithBinaryFiled struct {
	Content io.Reader `mandatory:"true" contributesTo:"body" encoding:"binary"`
}

func TestMarshalBinaryRequest(t *testing.T) {
	data := "some data in a file"
	buffer := bytes.NewBufferString(data)
	r := reqWithBinaryFiled{Content: ioutil.NopCloser(buffer)}
	httpRequest, err := MakeDefaultHTTPRequestWithTaggedStruct("PUT", "/obj", r)
	assert.NoError(t, err)
	all, err := ioutil.ReadAll(httpRequest.Body)
	assert.NoError(t, err)
	assert.Equal(t, data, string(all))
}

type structWithBinaryField struct {
	Content io.Reader `presentIn:"body" encoding:"binary"`
}

func TestUnmarshalResponse(t *testing.T) {
	data := "some data in a file"
	filename := writeTempFile(data)
	defer removeFileFn(filename)
	file, _ := os.Open(filename)
	header := http.Header{}
	r := http.Response{Header: header}
	r.Body = ioutil.NopCloser(file)
	s := structWithBinaryField{}
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	all, e := ioutil.ReadAll(s.Content)
	assert.NoError(t, e)
	assert.Equal(t, data, string(all))
}

type structWithHeaderCollections struct {
	Meta map[string]string `contributesTo:"header-collection" prefix:"meta-prefix-"`
}

func TestMarshalWithHeaderCollections(t *testing.T) {
	vals := make(map[string]string)
	vals["key1"] = "val1"
	vals["key2"] = "val2"
	s := structWithHeaderCollections{Meta: vals}

	request, err := MakeDefaultHTTPRequestWithTaggedStruct("GET", "/", s)
	assert.NoError(t, err)
	assert.Equal(t, s.Meta["key1"], request.Header.Get("meta-prefix-key1"))
	assert.Equal(t, s.Meta["key2"], request.Header.Get("meta-prefix-key2"))
}

func TestMarshalWithHeaderCollections_BadCollectionType(t *testing.T) {
	vals := make(map[string]int)
	vals["key1"] = 1
	s := struct {
		Meta map[string]int `contributesTo:"header-collection" prefix:"meta-prefix-"`
	}{Meta: vals}

	_, err := MakeDefaultHTTPRequestWithTaggedStruct("GET", "/", s)
	assert.Error(t, err)
}

type responseWithHC struct {
	Meta map[string]string `presentIn:"header-collection" prefix:"meta-prefix-"`
}

func TestUnMarshalWithHeaderCollections(t *testing.T) {
	header := http.Header{}
	s := responseWithHC{}
	header.Set("meta-prefix-key1", "val1")
	header.Set("meta-prefix-key2", "val2")
	r := http.Response{Header: header}
	err := UnmarshalResponse(&r, &s)
	assert.NoError(t, err)
	assert.Equal(t, s.Meta["key1"], r.Header.Get("val1"))
	assert.Equal(t, s.Meta["key2"], r.Header.Get("val2"))
}

type responseWithEmptyQP struct {
	Meta    string `contributesTo:"query" omitEmpty:"true" name:"meta"`
	QParam  string `contributesTo:"query" omitEmpty:"false" name:"qp"`
	QParam2 string `contributesTo:"query" name:"qp2"`
}

func TestEmptyQueryParam(t *testing.T) {
	s := responseWithEmptyQP{}
	r, err := MakeDefaultHTTPRequestWithTaggedStruct("GET", "/", s)
	assert.NoError(t, err)
	assert.Contains(t, r.URL.RawQuery, "qp2")
	assert.Contains(t, r.URL.RawQuery, "qp")
	assert.NotContains(t, r.URL.RawQuery, "meta")

}
