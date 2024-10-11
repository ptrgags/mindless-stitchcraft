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
