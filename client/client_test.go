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

func TestPerformRequest(t *testing.T) {
	// Create a test server to mock the API endpoint
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		expectedPath := "/some-endpoint"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
		}

		expectedParam1 := "value1"
		expectedParam2 := "value2"
		if gotParam1 := r.URL.Query().Get("Id"); gotParam1 != expectedParam1 {
			t.Errorf("Expected param1 %s, got %s", expectedParam1, gotParam1)
		}
		if gotParam2 := r.URL.Query().Get("Name"); gotParam2 != expectedParam2 {
			t.Errorf("Expected param2 %s, got %s", expectedParam2, gotParam2)
		}

		// Return a mock response
		response := `{"status": "success", "data": "some data"}`
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer testServer.Close()

	client := &Client{
		APIKey: "your-api-key",
	}

	url := testServer.URL + "/some-endpoint"
	params := map[string]string{"Id": "value1", "Name": "value2"}
	target := &SimpleResponse{}

	response := client.PerformRequest(url, params, target)
	if response == nil {
		t.Error("Expected a non-nil response")
	}
}

func TestCreatePath(t *testing.T) {
	expected := "https://inversionesyfinanzas.xyz/api/v1/basePath/?api_key=na&hola=hola"
	testClient := Client{APIKey: "na"}
	got := testClient.createPath("basePath", map[string]string{"hola": "hola"})
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
