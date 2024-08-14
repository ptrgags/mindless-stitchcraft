package patterns

import (
	"strings"
	"testing"
)

func TestSwapKnitsAndPurls(t *testing.T) {
	t.Run("invalid characters result in an error", func(t *testing.T) {
		cases := []struct {
			label string
			input string
		}{
			{"empty", ""},
			{"bad character", "-v🧶"},
		}

		for _, tc := range cases {
			swapped, err := SwapKnitsAndPurls(tc.input)
			if err == nil || !strings.Contains(err.Error(), "invalid stitches") {
				t.Errorf("(%v) Expected invalid stitches error, got (%v, %v)", tc.label, swapped, err)
			}
		}
	})

	t.Run("one type of stitch", func(t *testing.T) {
		cases := []struct {
			input    string
			expected string
		}{
			{"vvv", "---"},
			{"-----", "vvvvv"},
		}

		for _, tc := range cases {
			swapped, err := SwapKnitsAndPurls(tc.input)
			if swapped != tc.expected || err != nil {
				t.Errorf("Expected (%v, nil), got (%v, %v)", tc.expected, swapped, err)
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
			swapped, err := SwapKnitsAndPurls(tc.input)
			if swapped != tc.expected || err != nil {
				t.Errorf("Expected (%v, nil), got (%v, %v)", tc.expected, swapped, err)
			}
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
