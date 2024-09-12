package math

import (
	"strings"
	"testing"
)

func TestMakePermutation(t *testing.T) {
	t.Run("Zero length permutation returns error", func(t *testing.T) {
		perm, err := makePermutation([]uint{})

		expectedError := "must have at least one entry"
		if err == nil || !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error '%v', got (%v, %v)", expectedError, perm, err)
		}
	})

	t.Run("Out-of-bounds entry results in error", func(t *testing.T) {
		values := []uint{1, 2, 3, 4}

		perm, err := makePermutation(values)

		expectedError := "values must be in the range [0, 3]"
		if err == nil || !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error '%v', got (%v, %v)", expectedError, perm, err)
		}
	})

	t.Run("duplicate entry results in error", func(t *testing.T) {
		values := []uint{0, 0, 1, 2}

		perm, err := makePermutation(values)

		expectedError := "each entry must be listed exactly once"
		if err == nil || !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error '%v', got (%v, %v)", expectedError, perm, err)
		}
	})

	t.Run("Valid permutation does not produce error", func(t *testing.T) {
		values := []uint{0, 2, 1, 3}

		perm, err := makePermutation(values)

		if err != nil {
			t.Errorf(`Expected valid permutation, got (%v, %v)`, perm, err)
		}
	})
}
