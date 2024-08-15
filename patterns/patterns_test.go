package patterns

import (
	"fmt"
	"strings"
	"testing"
)

func TestValidateMotif(t *testing.T) {
	t.Run("empty motif returns error", func(t *testing.T) {
		err := validateMotif("")
		if err == nil || !strings.Contains(err.Error(), "must be a non-empty string") {
			t.Errorf("Expected non-empty error, got %v", err)
		}
	})

	t.Run("invalid motifs return error", func(t *testing.T) {
		err := validateMotif("🧶--")
		if err == nil || !strings.Contains(err.Error(), "must be a string of knits ('v') and purls ('-')") {
			t.Errorf("Expected non-empty error, got %v", err)
		}
	})

	t.Run("valid motif returns nil", func(t *testing.T) {
		err := validateMotif("v--v")
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})
}

func TestSwapKnitsAndPurls(t *testing.T) {

	t.Run("one type of stitch", func(t *testing.T) {
		cases := []struct {
			input    string
			expected string
		}{
			{"vvv", "---"},
			{"-----", "vvvvv"},
		}

		for _, tc := range cases {
			swapped := swapKnitsAndPurls(tc.input)
			if swapped != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, swapped)
			}
		}
	})

	t.Run("both knits and purls", func(t *testing.T) {
		cases := []struct {
			input    string
			expected string
		}{
			{"v---v", "-vvv-"},
			{"v--v-vv-", "-vv-v--v"},
		}

		for _, tc := range cases {
			swapped := swapKnitsAndPurls(tc.input)
			if swapped != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, swapped)
			}
		}
	})
}

func TestGenerateRawPattern(t *testing.T) {
	t.Fail()
}

func TestEnsureEvenRowCount(t *testing.T) {
	t.Fail()
}

func TestHandleReverseRows(t *testing.T) {
	t.Fail()
}

func checkRowsEqual(a []string, b []string) error {
	if len(a) != len(b) {
		return fmt.Errorf("Lengths don't match: len(a)=%v, len(b)=%v", a, b)
	}

	for i, aRow := range a {
		if aRow != b[i] {
			return fmt.Errorf("Row mismatch at position %d %v, %v a=%v, b=%v", i, aRow, b[i], a, b)
		}
	}

	return nil
}

func TestHandleRotate180(t *testing.T) {
	t.Run("input is not modified", func(t *testing.T) {
		input := []string{
			"v--",
			"---",
		}
		rotated := rotate180(input)
		if checkRowsEqual(rotated, input) == nil {
			t.Error("input was modified!")
		}
	})

	t.Run("Rotating nothing gives nothing", func(t *testing.T) {
		rotated := rotate180(nil)
		if rotated != nil {
			t.Errorf("Expected [], got %v", rotated)
		}
	})

	t.Run("Rotating a single row reverses the row", func(t *testing.T) {
		rotated := rotate180([]string{"vv---"})
		expected := []string{"---vv"}
		if err := checkRowsEqual(rotated, expected); err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("Rotating a single row reverses the row", func(t *testing.T) {
		rotated := rotate180([]string{"vv---"})
		expected := []string{"---vv"}
		if err := checkRowsEqual(rotated, expected); err != nil {
			t.Error(err.Error())
		}
	})
}

func TestGeneratePattern(t *testing.T) {
	t.Run("empty stitches string gives error", func(t *testing.T) {
		rows, err := GeneratePattern("", 5)
		if err == nil || !strings.Contains(err.Error(), "stitches must be non-empty") {
			t.Errorf("Expected non-empty stitches error, got (%v, %v)", rows, err)
		}
	})

	t.Run("invalid stitches gives error", func(t *testing.T) {
		rows, err := GeneratePattern("-🧶-", 5)
		if err == nil || !strings.Contains(err.Error(), "invalid stitches string") {
			t.Errorf("Expected invalid stitches error, got (%v, %v)", rows, err)
		}
	})

	t.Run("Invalid width gives error", func(t *testing.T) {
		widths := []int{0, -1}

		for _, width := range widths {
			rows, err := GeneratePattern("v--", width)
			if err == nil || !strings.Contains(err.Error(), "invalid width") {
				t.Errorf("(%v): Expected invalid width error, got (%v, %v)", width, rows, err)
			}
		}
	})

	t.Run("len(stitches) == width results in one row", func(t *testing.T) {
		rows, err := GeneratePattern("v-vv--", 6)
		if err != nil || len(rows) != 1 || rows[0] != "v-vv--" {
			t.Errorf(`Expected (["v-vv--"], nil), got (%v, %v)`, rows, err)
		}
	})

	t.Run("len(stitches) evenly divides width results in one row", func(t *testing.T) {
		rows, err := GeneratePattern("v--", 6)
		if err != nil || len(rows) != 1 || rows[0] != "v--v--" {
			t.Errorf(`Expected (["v--v--"], nil), got (%v, %v)`, rows, err)
		}
	})

	t.Run("len(stiches) < width, not coprime", func(t *testing.T) {
		t.Fail()
	})

	t.Run("len(stitches) < width, coprime", func(t *testing.T) {
		t.Fail()
	})

	t.Run("len(stitches) is a multiple of width", func(t *testing.T) {
		t.Fail()
	})

	t.Run("len(stitches) > width, not coprime", func(t *testing.T) {
		t.Fail()
	})

	t.Run("len(stitches) > width, not coprime", func(t *testing.T) {
		t.Fail()
	})
}
