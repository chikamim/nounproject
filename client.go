package nounproject

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/garyburd/go-oauth/oauth"
)

const (
	version     = "0.2"
	userAgent   = "go-nounproject/" + version
	httpTimeout = 30
)

var (
	baseURL = &url.URL{
		Scheme: "https",
		Host:   "api.thenounproject.com",
	}
)

// Client represents a Client object
type Client struct {
	httpClient  *http.Client
	oauthClient *oauth.Client
	BaseURL     *url.URL
	UserAgent   string
}

// NewClient returns a nounproject client
func NewClient(apiKey string, apiSecret string) *Client {
	oauthClient := &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  apiKey,
			Secret: apiSecret,
		},
	}
	timeout := time.Duration(httpTimeout * time.Second)
	httpClient := &http.Client{
		Timeout: timeout,
	}

	client := &Client{
		httpClient:  httpClient,
		oauthClient: oauthClient,
		BaseURL:     baseURL,
		UserAgent:   userAgent,
	}
	return client
}

// Get function that executes get request
func (c *Client) Get(u url.URL, result interface{}) (*http.Response, error) {
	u.Scheme = c.BaseURL.Scheme
	u.Host = c.BaseURL.Host
	req, err := http.NewRequest("GET", u.String(), nil)

	err = c.oauthClient.SetAuthorizationHeader(req.Header, nil, "GET", req.URL, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req, nil, result)
}

// Do function that executes http request
func (c *Client) Do(req *http.Request, body, result interface{}) (*http.Response, error) {
	if body != nil {
		bd, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		if req.Header == nil {
			req.Header = make(http.Header)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Body = ioutil.NopCloser(bytes.NewReader(bd))
		req.ContentLength = int64(len(bd))
	}

	res, err := c.httpClient.Do(req)
	if res.StatusCode == 404 {
		return res, fmt.Errorf("Not Found")
	}
	if res.StatusCode == 401 {
		return res, fmt.Errorf("Unauthorized")
	}
	if res.StatusCode == 403 {
		return res, fmt.Errorf("Forbidden")
	}
	if err != nil {
		return res, err
	}
	defer res.Body.Close()

	if result != nil {
		err := json.NewDecoder(res.Body).Decode(result)
		return res, err
	}
	return res, fmt.Errorf("No Result")
}
