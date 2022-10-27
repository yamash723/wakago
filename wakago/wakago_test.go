package wakago

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	assert.True(t, strings.HasSuffix(c.baseURL.String(), "/"))
	assert.Equal(t, c.baseURL.String(), defaultBaseURL)
	assert.Equal(t, c.UserAgent, defaultUserAgent)

	c2 := NewClient(nil)
	assert.Equal(t, c.client, c2.client)
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)

	testPath, expectedURL := "/foo", defaultBaseURL+"foo"
	testBody, expectBody := struct {
		Message string `json:"message"`
		Number  int    `json:"number"`
	}{"hello", 1}, `{"message":"hello","number":1}`+"\n"

	request, _ := c.NewRequest("GET", testPath, testBody)
	assert.Equal(t, request.URL.String(), expectedURL)

	userAgent := request.Header.Get("User-Agent")
	assert.Equal(t, userAgent, defaultUserAgent)
	assert.Contains(t, userAgent, Version)

	body, _ := ioutil.ReadAll(request.Body)
	assert.Equal(t, string(body), expectBody)
}

func TestNewRequest_defaultHeader(t *testing.T) {
	c := NewClient(nil)
	c.DefaultHeader = &http.Header{"HEADER-KEY": []string{"HEADER-VALUE"}}

	request, _ := c.NewRequest("GET", "/foo", nil)
	assert.Equal(t, []string{"HEADER-VALUE"}, request.Header["HEADER-KEY"])
}

func TestNewRequest_emptyBody(t *testing.T) {
	c := NewClient(nil)
	_, err := c.NewRequest("GET", "/foo", nil)
	assert.Nil(t, err)
}

func TestNewRequest_invalidJSON(t *testing.T) {
	c := NewClient(nil)

	type T struct {
		A map[interface{}]interface{}
	}

	_, err := c.NewRequest("GET", "/foo", &T{})
	assert.NotNil(t, err)

	assert.NotNil(t, err.(*json.UnsupportedTypeError))
}

func TestNewRequest_badMethod(t *testing.T) {
	c := NewClient(nil)
	_, err := c.NewRequest("BOGUS\nMETHOD", "/foo", nil)
	assert.NotNil(t, err)
}

func TestNewRequest_emptyUserAgent(t *testing.T) {
	c := NewClient(nil)
	c.UserAgent = ""

	request, err := c.NewRequest("GET", "/foo", nil)
	assert.Nil(t, err)

	_, isContains := request.Header["User-Agent"]
	assert.False(t, isContains)
}

func TestDo(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://wakatime.com/api/v1/foo",
		httpmock.NewStringResponder(200, `{"success":true}`))

	type foo struct {
		Success bool `json:"success"`
	}

	c := NewClient(nil)
	request, _ := c.NewRequest("GET", "/foo", nil)

	ctx := context.Background()
	v := new(foo)

	response, err := c.Do(ctx, request, v)
	assert.Nil(t, err)
	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, &foo{true}, v)
}

func TestDo_nilContext(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://wakatime.com/api/v1/foo",
		httpmock.NewStringResponder(200, `{"success":true}`))

	type foo struct {
		Success bool `json:"success"`
	}

	c := NewClient(nil)
	request, _ := c.NewRequest("GET", "/foo", nil)

	v := new(foo)

	_, err := c.Do(nil, request, v)
	assert.NotNil(t, err)
}

func TestDo_httpError(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://wakatime.com/api/v1/foo",
		httpmock.NewStringResponder(400, `{"success":false}`))

	type foo struct {
		Success bool `json:"success"`
	}

	c := NewClient(nil)
	request, _ := c.NewRequest("GET", "/foo", nil)

	ctx := context.Background()
	v := new(foo)

	response, err := c.Do(ctx, request, v)
	assert.NotNil(t, err)
	assert.Equal(t, 400, response.StatusCode)
}
