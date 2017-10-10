package common

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Test data structures, avoid import cycle
type UpdateUserDetails struct {
	Description string `mandatory:"false" json:"description,omitempty"`
}

type ListCompartmentsRequest struct {
	CompartmentID string `mandatory:"true" contributesTo:"query" name:"compartmentId"`
	Page          string `mandatory:"false" contributesTo:"query" name:"page"`
	Limit         int32  `mandatory:"false" contributesTo:"query" name:"limit"`
}

type UpdateUserRequest struct {
	UserID            string            `mandatory:"true" contributesTo:"path" name:"userId"`
	UpdateUserDetails UpdateUserDetails `contributesTo:"body"`
	IfMatch           string            `mandatory:"false" contributesTo:"header" name:"if-match"`
}

type CreateApiKeyDetails struct {
	Key string `mandatory:"true" json:"key,omitempty"`
}

type UploadApiKeyRequest struct {
	UserID              string              `mandatory:"true" contributesTo:"path" name:"userId"`
	CreateApiKeyDetails CreateApiKeyDetails `contributesTo:"body"`
	OpcRetryToken       string              `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestHttpMarshallerInvalidStruct(t *testing.T) {
	request := http.Request{}
	err := HttpRequestMarshaller("asdf", &request)
	assert.Error(t, err, nil)
}

func TestHttpRequestMarshallerQuery(t *testing.T) {
	s := ListCompartmentsRequest{CompartmentID: "ocid1", Page: "p", Limit: 23}
	request := &http.Request{}
	HttpRequestMarshaller(s, request)
	query := request.URL.Query()
	assert.True(t, query.Get("compartmentId") == "ocid1")
	assert.True(t, query.Get("page") == "p")
	assert.True(t, query.Get("limit") == "23")
}

func TestMakeDefault(t *testing.T) {
	r := MakeDefaultHttpRequest(http.MethodPost, "/one/two")
	assert.Equal(t, r.Header.Get("Content-Type"), "application/json")
	assert.NotEmpty(t, r.Header.Get("Date"))
	assert.NotEmpty(t, r.Header.Get("Opc-Client-Info"))
}

func TestHttpMarshallerSimpleHeader(t *testing.T) {
	s := UpdateUserRequest{UserID: "id1", IfMatch: "n=as", UpdateUserDetails: UpdateUserDetails{Description: "name of"}}
	request := MakeDefaultHttpRequest(http.MethodPost, "/random")
	HttpRequestMarshaller(s, &request)
	header := request.Header
	assert.True(t, header.Get("if-match") == "n=as")
}

func TestHttpMarshallerSimpleStruct(t *testing.T) {
	s := UploadApiKeyRequest{UserID: "111", OpcRetryToken: "token", CreateApiKeyDetails: CreateApiKeyDetails{Key: "thekey"}}
	request := MakeDefaultHttpRequest(http.MethodPost, "/random")
	HttpRequestMarshaller(s, &request)
	assert.True(t, strings.Contains(request.URL.Path, "111"))
}
func TestHttpMarshallerSimpleBody(t *testing.T) {
	desc := "theDescription"
	s := UpdateUserRequest{UserID: "id1", IfMatch: "n=as", UpdateUserDetails: UpdateUserDetails{Description: desc}}
	request := MakeDefaultHttpRequest(http.MethodPost, "/random")
	HttpRequestMarshaller(s, &request)
	body, _ := ioutil.ReadAll(request.Body)
	var content map[string]string
	json.Unmarshal(body, &content)
	assert.Contains(t, content, "description")
	if val, ok := content["description"]; !ok || val != desc {
		assert.Fail(t, "Should contain: "+desc)
	}
}

func TestHttpMarshalerAll(t *testing.T) {
	desc := "theDescription"
	s := struct {
		Id      string            `contributesTo:"path"`
		Name    string            `contributesTo:"query" name:"name"`
		When    time.Time         `contributesTo:"query" name:"when"`
		Income  float32           `contributesTo:"query" name:"income"`
		Male    bool              `contributesTo:"header" name:"male"`
		Details UpdateUserDetails `contributesTo:"body"`
	}{
		"101", "tapir", time.Now(), 3.23, true, UpdateUserDetails{Description: desc},
	}
	request := MakeDefaultHttpRequest(http.MethodPost, "/")
	HttpRequestMarshaller(s, &request)
	var content map[string]string
	body, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(body, &content)
	when := s.When.Format(time.RFC3339)
	assert.True(t, request.URL.Path == "/101")
	assert.True(t, request.URL.Query().Get("name") == s.Name)
	assert.True(t, request.URL.Query().Get("income") == strconv.FormatFloat(float64(s.Income), 'f', 6, 32))
	assert.True(t, request.URL.Query().Get("when") == when)
	assert.Contains(t, content, "description")
	if val, ok := content["description"]; !ok || val != desc {
		assert.Fail(t, "Should contain: "+desc)
	}
}

func TestHttpMarshalerPointers(t *testing.T) {

	var n *string = new(string)
	*n = "theName"
	s := struct {
		Name *string `contributesTo:"query" name:"name"`
	}{
		n,
	}
	request := MakeDefaultHttpRequest(http.MethodPost, "/random")
	HttpRequestMarshaller(s, &request)
	assert.NotNil(t, request)
	assert.True(t, request.URL.Query().Get("name") == *s.Name)
}

func TestHttpMarshalerUntaggedFields(t *testing.T) {
	s := struct {
		Name  string `contributesTo:"query" name:"name"`
		AList []string
		AMap  map[string]int
		UpdateUserDetails
	}{
		"theName", []string{"a", "b"}, map[string]int{"a": 1, "b": 2},
		UpdateUserDetails{Description: "n"},
	}
	request := &http.Request{}
	e := HttpRequestMarshaller(s, request)
	assert.NoError(t, e)
	assert.NotNil(t, request)
	assert.True(t, request.URL.Query().Get("name") == s.Name)
}
func TestHttpMarshalerPathTemplate(t *testing.T) {
	urlTemplate := "/name/{userId}/aaa"
	s := UploadApiKeyRequest{UserID: "111", OpcRetryToken: "token", CreateApiKeyDetails: CreateApiKeyDetails{Key: "thekey"}}
	request := MakeDefaultHttpRequest(http.MethodPost, urlTemplate)
	e := HttpRequestMarshaller(s, &request)
	assert.NoError(t, e)
	assert.Equal(t, "/name/111/aaa", request.URL.Path)
}

func TestHttpMarshalerFunnyTags(t *testing.T) {
	s := struct {
		Name  string `contributesTo:"quer" name:"name"`
		AList []string
		AMap  map[string]int
		UpdateUserDetails
	}{
		"theName", []string{"a", "b"}, map[string]int{"a": 1, "b": 2},
		UpdateUserDetails{Description: "n"},
	}
	request := &http.Request{}
	e := HttpRequestMarshaller(s, request)
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
		Name              string `contributesTo:"query" name:"name"`
		UpdateUserDetails `contributesTo:"query" name:"str"`
	}{
		"theName", UpdateUserDetails{Description: "a"},
	}
	var n *string = new(string)
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
		e := HttpRequestMarshaller(l, request)
		Debugln(e)
		assert.Error(t, e)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// BaseClient
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestBaseClient_prepareRequest(t *testing.T) {
	r := MakeDefaultHttpRequest(http.MethodPost, "/random")
	c := NewClient()
	c.ApiVersion = "v1"
	e := c.prepareRequest(&r)
	assert.NoError(t, e)
	assert.Equal(t, "/v1/random", r.URL.Path)

}
