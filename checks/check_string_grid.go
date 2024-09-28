package checks

import (
	"testing"
)

func CheckStringGridShape(t *testing.T, rows []string, expectedWidth int, expectedHeight int) {
	if len(rows) != expectedHeight {
		t.Errorf("Incorrect number of rows. Expected %v, got %v, rows=%v", expectedHeight, len(rows), rows)
	}

	for _, row := range rows {
		if len(row) != expectedWidth {
			t.Errorf("Incorrect number of cols. Expected %v, got %v. rows=%v", expectedWidth, len(rows), rows)
		}
	}
}

func CheckStringGridsEqual(t *testing.T, actual []string, expected []string) error {
	if len(actual) != len(expected) {
		t.Errorf("Lengths don't match: len(actual)=%v, len(expected)=%v, actual=%v, expected=%v", len(actual), len(expected), actual, expected)
	}

	for i, actualRow := range actual {
		if actualRow != expected[i] {
			t.Errorf("Row mismatch at position %d: %v, %v actual=%v, expected=%v", i, actualRow, expected[i], actual, expected)
		}
	}

	return nil
}
