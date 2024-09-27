package checks

import (
	"strings"
	"testing"
)

func CheckHasError(t *testing.T, result any, err error, expectedMessage string) {
	if err == nil || !strings.Contains(err.Error(), expectedMessage) {
		t.Errorf("Expected error %v, got (%v, %v)", expectedMessage, result, err)
	}
}

func CheckHasNoError(t *testing.T, result any, err error) {
	if err != nil {
		t.Errorf("Expected no error, got (%v, %v)", result, err)
	}
}
