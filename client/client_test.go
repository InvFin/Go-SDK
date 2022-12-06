package client_test

import (
	"invfinsdk/client"
	"testing"
)

func TestPerformRequest(t *testing.T) {
	expected := "https://inversionesyfinanzas.xyz/api/v1/basePath/?api_key=na&hola=hola"
	testClient := client.Client{APIKey: "na"}
	got := testClient.PerformRequest("basePath", map[string]string{"hola": "hola"}, map[string]string{})
	if got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
