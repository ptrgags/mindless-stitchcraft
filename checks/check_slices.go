package checks

import "testing"

func CheckSliceEmpty[T any](t *testing.T, slice []T) {
	if len(slice) != 0 {
		t.Errorf("Expected empty slice, got %v", slice)
	}
}

func CheckSlicesEqual[T comparable](t *testing.T, actual []T, expected []T) {
	if len(actual) != len(expected) {
		t.Errorf("Lengths don't match: len(actual)=%v, len(expected)=%v, actual=%v, expected=%v", len(actual), len(expected), actual, expected)
		return
	}

	for i, actualValue := range actual {
		if actualValue != expected[i] {
			t.Errorf("Value mismatch at position %d: %v, %v actual=%v, expected=%v", i, actualValue, expected[i], actual, expected)
		}
	}
}

func CheckNestedSlicesEqual[T comparable](t *testing.T, actual [][]T, expected [][]T) {
	if len(actual) != len(expected) {
		t.Errorf("Row counts don't match: rows(actual)=%v, rows(expected)=%v, actual=%v, expected=%v", len(actual), len(expected), actual, expected)
		return
	}

	for i, actualRow := range actual {
		expectedRow := expected[i]
		if len(actualRow) != len(expectedRow) {
			t.Errorf("Row %d: Column counts don't match: cols(actual)=%v, cols(expected)=%v, actual=%v, expected=%v", i, len(actualRow), len(expectedRow), actual, expected)
			continue
		}

		for j, actualValue := range actualRow {
			if actualValue != expectedRow[j] {
				t.Errorf("Row %d, col %d: value mismatch: %v != %v, actual=%v, expected=%v", i, j, actualValue, expectedRow[i], actual, expected)
			}
		}
	}
}
