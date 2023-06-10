package api_test

import (
	"testing"

	"github.com/InvFin/Go-SDK/api"
)

func TestNewAPI(t *testing.T) {
	got, err := api.NewAPI("na")
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if got == nil {
		t.Error("Expected API instance, got nil")
	}
}

func TestEmptyNewAPI(t *testing.T) {
	api, err := api.NewAPI("")
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if api != nil {
		t.Error("Expected nil API instance, got non-nil")
	}
}
