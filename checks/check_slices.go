package checks

import "testing"

func CheckSlicesEqual[T comparable](t *testing.T, actual []T, expected []T) {
	if len(actual) != len(expected) {
		t.Errorf("Lengths don't match: len(actual)=%v, len(expected)=%v, actual=%v, expected=%v", len(actual), len(expected), actual, expected)
	}

	for i, actualValue := range actual {
		if actualValue != expected[i] {
			t.Errorf("Value mismatch at position %d: %v, %v actual=%v, expected=%v", i, actualValue, expected[i], actual, expected)
		}
	}
}
