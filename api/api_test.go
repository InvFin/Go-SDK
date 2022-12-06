package api_test

import (
	"invfinsdk/api"
	"testing"
)

func TestNewAPI(t *testing.T) {
	expected, _ := api.NewAPI(client.Config{APIKey: "na"})
	got := client.Client{APIKey: "na"}
	if got != expected {
		t.Fatal()
	}
}
