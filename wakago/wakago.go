package wakago

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	Version = "0.0.1"

	defaultBaseURL   = "https://wakatime.com/api/v1/"
	defaultUserAgent = "wakago" + "/" + Version
)

type Client struct {
	client  *http.Client
	baseURL *url.URL

	UserAgent     string
	DefaultHeader *http.Header

	AllTimeSinceTodayService *AllTimeSinceTodayService
	MetaService              *MetaService
	EditorsService           *EditorsService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: &http.Client{}, baseURL: baseURL, UserAgent: defaultUserAgent}
	c.AllTimeSinceTodayService = &AllTimeSinceTodayService{client: c}
	c.MetaService = &MetaService{client: c}
	c.EditorsService = &EditorsService{client: c}

	return c
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	url := c.baseURL.JoinPath(path)

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, url.String(), buf)

	if err != nil {
		return nil, err
	}

	if c.DefaultHeader != nil {
		request.Header = *c.DefaultHeader
	}

	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	if c.UserAgent != "" {
		request.Header.Set("User-Agent", c.UserAgent)
	}

	return request, nil
}

func (c *Client) Do(ctx context.Context, request *http.Request, v interface{}) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("context is nil")
	}

	request = request.WithContext(ctx)

	response, err := c.client.Do(request)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	if err := CheckHttpStatusCode(response.StatusCode); err != nil {
		return response, err
	}

	if v != nil {
		err = json.NewDecoder(response.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}

	return response, err
}

// HTTP response codes
//
// 200 - Ok: The request has succeeded.
// 201 - Created: The request has been fulfilled and resulted in a new resource being created.
// 202 - Accepted: The request has been accepted for processing, but the processing has not been completed. The stats resource may return this code.
// 400 - Bad Request: The request is invalid. Check error message and try again.
// 401 - Unauthorized: The request requires authentication, or your authentication was invalid.
// 403 - Forbidden: You are authenticated, but do not have permission to access the resource.
// 404 - Not Found: The resource does not exist.
// 429 - Too Many Requests: You are being rate limited, try making fewer than 10 requests per second on average over any 5 minute period.
// 500 - Server Error: Service unavailable, try again later.
//
// refer : https://wakatime.com/developers#introduction
func CheckHttpStatusCode(statusCode int) error {
	if statusCode >= 200 && statusCode <= 399 {
		return nil
	}

	return errors.New(fmt.Sprintf("Return error HTTP response code : %v", statusCode))
}
