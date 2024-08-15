package patterns

import (
	"fmt"
	"testing"
)

/*
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
	t.Run("no rows returns an empty slice", func(t *testing.T) {
		result := ensureEvenRowCount([]string{})
		if err := checkRowsEqual(result, []string{}); err != nil {
			t.Error(err)
		}
	})

	t.Run("Even input rows results in the same rows", func(t *testing.T) {
		input := []string{
			"v---v-",
			"v-vvvv",
		}
		result := ensureEvenRowCount(input)
		if err := checkRowsEqual(result, input); err != nil {
			t.Error(err)
		}
	})

	t.Run("Odd input rows results in doubled output", func(t *testing.T) {
		input := []string{
			"v---v-",
			"v-vvvv",
			"------",
		}
		expected := []string{
			"v---v-",
			"v-vvvv",
			"------",
			"v---v-",
			"v-vvvv",
			"------",
		}

		result := ensureEvenRowCount(input)

		if err := checkRowsEqual(result, expected); err != nil {
			t.Error(err)
		}
	})

	t.Run("Input slice is not modified", func(t *testing.T) {
		t.Fail()
	})
}

func TestHandleReverseRows(t *testing.T) {
	t.Fail()
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
}*/

func checkRowsEqual(a []string, b []string) error {
	if len(a) != len(b) {
		return fmt.Errorf("Lengths don't match: len(a)=%v, len(b)=%v, a=%v, b=%v", len(a), len(b), a, b)
	}

	for i, aRow := range a {
		if aRow != b[i] {
			return fmt.Errorf("Row mismatch at position %d %v, %v a=%v, b=%v", i, aRow, b[i], a, b)
		}
	}

	return nil
}

func checkShape(rows []string, expectedWidth int, expectedHeight int) error {
	if len(rows) != expectedHeight {
		return fmt.Errorf("Incorrect number of rows. Expected %v, got %v, rows=%v", expectedHeight, len(rows), rows)
	}

	for _, row := range rows {
		if len(row) != expectedWidth {
			return fmt.Errorf("Incorrect number of cols. Expected %v, got %v. rows=%v", expectedWidth, len(rows), rows)
		}
	}

	return nil
}

func TestGeneratePrintablePattern(t *testing.T) {
	t.Run("invalid input", func(t *testing.T) {
		t.Fail()
	})

	t.Run("generated pattern is the correct shape", func(t *testing.T) {
		// The resulting pattern will always be:
		// width: fabricWidth
		// height: len(motif) / gcd(len(motif), fabricWidth)
		//    or twice this if there was an odd number of rows
		cases := []struct {
			label          string
			motif          string
			fabricWidth    int
			expectedHeight int
		}{
			{"len(motif) < fabricWidth, coprime widths", "v---", 5, 4},
			// height = 6 / gcd(6, 9) = 6 / 3 = 2
			{"len(motif) < fabricWidth, noncoprime widths", "---vv-", 9, 2},
			{"len(motif) > fabricWidth, coprime widths", "vvvv----", 3, 8},
			// height = 4 / gcd(4, 2) = 4 / 2 = 2
			{"len(motif) > fabricWidth, noncoprime widths", "----", 2, 2},
		}

		for _, tc := range cases {
			rows, err := GeneratePrintablePattern(tc.motif, tc.fabricWidth)
			if err != nil {
				t.Errorf("Expected no errors, got (%v, %v)", rows, err)
			}

			shapeErr := checkShape(rows, tc.fabricWidth, tc.expectedHeight)
			if shapeErr != nil {
				t.Errorf("%v: %v", tc.label, shapeErr.Error())
			}
		}
	})

	t.Run("generated pattern is doubled to ensure even length", func(t *testing.T) {
		// The resulting pattern will always be:
		// width: fabricWidth
		// height: len(motif) / gcd(len(motif), fabricWidth)
		//    or twice this if there was an odd number of rows
		cases := []struct {
			label          string
			motif          string
			fabricWidth    int
			expectedHeight int
		}{
			// height 1 -> 2
			{"len(motif) == fabricWidth", "v--", 3, 2},
			// height 3 -> 6
			{"len(motif) < fabricWidth, coprime widths", "v--", 5, 6},
			// height = 5 / gcd(5, 10) = 5 / 5 = 1 -> 2
			{"len(motif) < fabricWidth, noncoprime widths", "v--v-", 10, 2},
			// height 5 -> 10
			{"len(motif) > fabricWidth, coprime widths", "-vvv-", 3, 10},
			// height = 9 / gcd(9, 3) = 9 / 3 = 3 -> 6
			{"len(motif) > fabricWidth, noncoprime widths", "--v---vvv", 3, 6},
		}

		for _, tc := range cases {
			rows, err := GeneratePrintablePattern(tc.motif, tc.fabricWidth)
			if err != nil {
				t.Errorf("Expected no errors, got (%v, %v)", rows, err)
			}

			shapeErr := checkShape(rows, tc.fabricWidth, tc.expectedHeight)
			if shapeErr != nil {
				t.Errorf("%v: %v", tc.label, shapeErr.Error())
			}
		}
	})

	t.Run("generated pattern shows right side of fabric", func(t *testing.T) {
		cases := []struct {
			label        string
			motif        string
			fabricWidth  int
			expectedRows []string
		}{
			{"len(motif) < fabricWidth, coprime widths", "v---", 5, []string{
				"v---",
			}},
			// height = 6 / gcd(6, 9) = 6 / 3 = 2
			{"len(motif) < fabricWidth, noncoprime widths", "---vv-", 9, []string{}},
			{"len(motif) > fabricWidth, coprime widths", "vvvv----", 3, []string{}},
			// height = 4 / gcd(4, 2) = 4 / 2 = 2
			{"len(motif) > fabricWidth, noncoprime widths", "----", 2, []string{}},
		}
		for _, tc := range cases {
			rows, err := GeneratePrintablePattern(tc.motif, tc.fabricWidth)
			if err != nil {
				t.Errorf("Expected no errors, got (%v, %v)", rows, err)
			}

			equalsErr := checkRowsEqual(rows, tc.expectedRows)
			if equalsErr != nil {
				t.Errorf("%v: %v", tc.label, equalsErr.Error())
			}
		}
	})

	t.Run("generated pattern shows right side of fabric (doubled cases)", func(t *testing.T) {
		cases := []struct {
			label        string
			motif        string
			fabricWidth  int
			expectedRows []string
		}{
			{"len(motif) == fabricWidth", "v--", 3, []string{
				"v--",
				"vv-",
			}},
			{"len(motif) < fabricWidth, coprime widths", "v--", 5, []string{
				"v--v-",
				"-vv-v",
				"--v--",
				"v--v-",
				"-v--v",
				"--v--",
			}},
			{"len(motif) < fabricWidth, noncoprime widths", "v--v-", 10, []string{}},
			{"len(motif) > fabricWidth, coprime widths", "-vvv-", 3, []string{}},
			{"len(motif) > fabricWidth, noncoprime widths", "--v---vvv", 3, []string{}},
		}

		for _, tc := range cases {
			rows, err := GeneratePrintablePattern(tc.motif, tc.fabricWidth)
			if err != nil {
				t.Errorf("Expected no errors, got (%v, %v)", rows, err)
			}

			equalsErr := checkRowsEqual(rows, tc.expectedRows)
			if equalsErr != nil {
				t.Errorf("%v: %v", tc.label, equalsErr.Error())
			}
		}
	})
}
