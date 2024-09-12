package math

import (
	"strings"
	"testing"
)

func TestMakePermutation(t *testing.T) {
	t.Run("Zero length permutation returns error", func(t *testing.T) {
		perm, err := MakePermutation([]uint{})

		expectedError := "must have at least one entry"
		if err == nil || !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error '%v', got (%v, %v)", expectedError, perm, err)
		}
	})

	t.Run("Out-of-bounds entry results in error", func(t *testing.T) {
		values := []uint{1, 2, 3, 4}

		perm, err := MakePermutation(values)

		expectedError := "values must be in the range [0, 3]"
		if err == nil || !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error '%v', got (%v, %v)", expectedError, perm, err)
		}
	})

	t.Run("duplicate entry results in error", func(t *testing.T) {
		values := []uint{0, 0, 1, 2}

		perm, err := MakePermutation(values)

		expectedError := "each entry must be listed exactly once"
		if err == nil || !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error '%v', got (%v, %v)", expectedError, perm, err)
		}
	})

	t.Run("Valid permutation does not produce error", func(t *testing.T) {
		values := []uint{0, 2, 1, 3}

		perm, err := MakePermutation(values)

		if err != nil {
			t.Errorf("Expected valid permutation, got (%v, %v)", perm, err)
		}
	})
}

func TestApplyPermutation(t *testing.T) {
	t.Run("out-of-range value is unmodified", func(t *testing.T) {
		perm, _ := MakePermutation([]uint{0, 2, 1, 3})
		var outOfRange uint = 6

		result := Apply(perm, outOfRange)

		if result != outOfRange {
			t.Errorf("Expected %v, got %v", outOfRange, result)
		}
	})

	t.Run("in-range value is cycled correctly", func(t *testing.T) {
		perm, _ := MakePermutation([]uint{0, 2, 1, 3})
		var shouldBeScrambled uint = 1

		result := Apply(perm, shouldBeScrambled)

		var expected uint = 2
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("in-range value is fixed correctly", func(t *testing.T) {
		// This permutation only swaps 1 and 2, so 0 and 3 should remain
		// fixed.
		perm, _ := MakePermutation([]uint{0, 2, 1, 3})
		var shouldBeFixed uint = 3

		result := Apply(perm, shouldBeFixed)

		if result != shouldBeFixed {
			t.Errorf("Expected %v, got %v", shouldBeFixed, result)
		}
	})
}

func TestComposePermutations(t *testing.T) {
	t.Run("mismatched permutation lengths results in error", func(t *testing.T) {
		perm3, _ := MakePermutation([]uint{0, 2, 1})
		perm4, _ := MakePermutation([]uint{1, 3, 0, 2})

		result, err := Compose(perm4, perm3)

		expectedError := "permutations must have the same length"
		if err == nil || !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error '%v', got (%v, %v)", expectedError, result, err)
		}
	})

	t.Run("composing valid permutations do not produce error", func(t *testing.T) {
		permA, _ := MakePermutation([]uint{1, 2, 0, 3})
		permB, _ := MakePermutation([]uint{0, 2, 3, 1})

		result, err := Compose(permA, permB)

		if err != nil {
			t.Errorf("Expected no error, got (%v, %v)", result, err)
		}
	})

	t.Run("permutations are composed from right to left", func(t *testing.T) {
		permA, _ := MakePermutation([]uint{1, 2, 0, 3}) // (0 1 2)
		permB, _ := MakePermutation([]uint{0, 2, 3, 1}) // (1 2 3)

		result, _ := Compose(permA, permB)

		// permutations are applied from right to left, so we have
		// (0 1 2)(1 2 3) = (0 1)(2 3) which is  [1, 0, 3, 2]
		// if it were computed from left to right, you'd get:
		// (0 2)(1 3) which is not what we want here.
		expectedPerm, _ := MakePermutation([]uint{1, 0, 3, 2})
		if !Equals(result, expectedPerm) {
			t.Errorf("Expected %v, got %v", expectedPerm, result)
		}
	})
}

func TestPermutationEquals(t *testing.T) {
	t.Run("different permutation lengths returns false", func(t *testing.T) {
		perm3, _ := MakePermutation([]uint{1, 2, 0})
		perm4, _ := MakePermutation([]uint{2, 1, 0, 3})

		result := Equals(perm3, perm4)

		if result == true {
			t.Errorf("Expected false, got %v", result)
		}
	})

	t.Run("same permutation returns true", func(t *testing.T) {
		permA, _ := MakePermutation([]uint{1, 2, 0})
		permB, _ := MakePermutation([]uint{1, 2, 0})

		result := Equals(permA, permB)

		if result == false {
			t.Errorf("Expected true, got %v", result)
		}
	})

	t.Run("different permutation returns false", func(t *testing.T) {
		permA, _ := MakePermutation([]uint{1, 2, 0})
		permB, _ := MakePermutation([]uint{1, 0, 2})

		result := Equals(permA, permB)

		if result == true {
			t.Errorf("Expected false, got %v", result)
		}
	})
}
