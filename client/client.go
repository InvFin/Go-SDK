package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"
)

const (
	baseAPIPath string = "https://inversionesyfinanzas.xyz/api"
	apiVersion  string = "v1"
	baseHeader  string = "invfinsdk-Go"
)

type Client struct {
	APIKey string
}

func (c *Client) createPath(path string, params map[string]string) string {
	basePath := fmt.Sprintf("%s/%s/%s/", baseAPIPath, apiVersion, path)
	payload := url.Values{}
	payload.Add("api_key", c.APIKey)
	for key, value := range params {
		payload.Add(key, value)
	}
	return basePath + payload.Encode()
}

func (c *Client) buildRequest(path string) *http.Request {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		panic(err)
	}
	userAgent := fmt.Sprintf("%s-%s", baseHeader, apiVersion)
	req.Header.Add("User-Agent", userAgent)
	return req
}

func (c *Client) getCleanResponse(req *http.Request, target interface{}) interface{} {
	httpClient := http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func (c *Client) PerformRequest(url string, params map[string]string, target interface{}) interface{} {
	path := c.createPath(url, params)
	req := c.buildRequest(path)
	return c.getCleanResponse(req, target)
}

func TestcreatePath(t *testing.T) {
	expected := "https://inversionesyfinanzas.xyz/api/v1/basePath/?api_key=na&hola=hola"
	testClient := Client{APIKey: "na"}
	got := testClient.createPath("basePath", map[string]string{"hola": "hola"})
	if got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
