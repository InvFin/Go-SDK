package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

type ErrorResponse struct {
	StatusCode int
	Detail     string
}

func (c *Client) PerformRequest(url string, params map[string]string, target interface{}) interface{} {
	return c.getCleanResponse(c.buildRequest(c.createPath(url, params)), target)
}

func (c *Client) createPath(path string, params map[string]string) string {
	basePath := fmt.Sprintf("%s/%s/%s/?", baseAPIPath, apiVersion, path)
	payload := url.Values{}
	payload.Add("api_key", c.APIKey)
	if params != nil {
		for key, value := range params {
			payload.Add(key, value)
		}
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
	jsonDecoder := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK {
		errorResponse := ErrorResponse{
			StatusCode: resp.StatusCode,
		}
		if err := jsonDecoder.Decode(&errorResponse); err != nil {
			panic(err)
		}
		return errorResponse
	}
	if err := jsonDecoder.Decode(target); err != nil {
		panic(err)
	}
	return target
}
