package services_test

import (
	"testing"

	"github.com/braydend/my-family/services"
)

func TestNewLoggerService(t *testing.T) {
	logger := services.NewLogger()
	if logger == nil {
		t.Error("Expected logger to be created")
	}
}
