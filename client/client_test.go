package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SimpleResponse struct {
	Id   int
	Name string
}

func TestCreatePath(t *testing.T) {
	expected := "https://inversionesyfinanzas.xyz/api/v1/basePath/?api_key=na&hola=hola"
	testClient := Client{APIKey: "na"}
	got := testClient.createPath("basePath", map[string]string{"hola": "hola"})
	if got != expected {
		t.Errorf("expected %s and got %s", expected, got)
	}
}

func TestCreatePathEmpty(t *testing.T) {
	expected := "https://inversionesyfinanzas.xyz/api/v1/basePath/?api_key=na"
	testClient := Client{APIKey: "na"}
	got := testClient.createPath("basePath", nil)
	if got != expected {
		t.Errorf("expected %s and got %s", expected, got)
	}
}

func TestBuildRequest(t *testing.T) {
	client := &Client{
		APIKey: "your-api-key",
	}

	path := "https://example.com/api"
	req := client.buildRequest(path)

	if req.Method != "GET" {
		t.Errorf("Expected GET method, got: %s", req.Method)
	}

	if req.URL.String() != path {
		t.Errorf("Expected URL: %s, got: %s", path, req.URL.String())
	}

	expectedUserAgent := fmt.Sprintf("%s-%s", "invfinsdk-Go", "v1")
	actualUserAgent := req.Header.Get("User-Agent")
	if actualUserAgent != expectedUserAgent {
		t.Errorf("Expected User-Agent: %s, got: %s", expectedUserAgent, actualUserAgent)
	}
}

func TestGetCleanResponse(t *testing.T) {
	// Create a mock HTTP server with the desired response
	mockResponse := `{"id": 1, "name": "John Doe"}`
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	client := &Client{
		APIKey: "your-api-key",
	}

	req, err := http.NewRequest("GET", mockServer.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	var target SimpleResponse
	result := client.getCleanResponse(req, &target)

	if result != nil {
		t.Errorf("Expected nil result, got: %v", result)
	}

	expectedTarget := SimpleResponse{Id: 1, Name: "John Doe"}
	if target != expectedTarget {
		t.Errorf("Expected target: %+v, got: %+v", expectedTarget, target)
	}
}
