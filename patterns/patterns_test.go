package patterns

import (
	"strings"
	"testing"
)

func TestSwapKnitsAndPurls(t *testing.T) {

	t.Run("empty input is unchanged", func(t *testing.T) {
		swapped, err := SwapKnitsAndPurls("")
		if swapped != "" || err != nil {
			t.Errorf("Expected ('', nil), got (%v, %v)", swapped, err)
		}
	})

	t.Run("invalid characters result in an error", func(t *testing.T) {
		// Knits are represented as v in this program...
		badChars := "🧶vvv"
		swapped, err := SwapKnitsAndPurls(badChars)
		if err == nil || !strings.Contains(err.Error(), "invalid stitches") {
			t.Errorf("Expected invalid stitches error, got (%v, %v)", swapped, err)
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
