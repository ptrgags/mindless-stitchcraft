package knitting

import (
	"testing"

	"github.com/ptrgags/mindless-stitchcraft/checks"
)

func TestKnitStitchToRune(t *testing.T) {
	t.Run("Returns v for Knit", func(t *testing.T) {
		result := Knit.ToRune()

		if result != 'v' {
			t.Errorf("Expected v, got %v", result)
		}
	})

	t.Run("Returns - for Purl", func(t *testing.T) {
		result := Purl.ToRune()

		if result != '-' {
			t.Errorf("Expected -, got %v", result)
		}
	})
}

func TestKnitStitchSwap(t *testing.T) {
	t.Run("Returns Purl for Knit", func(t *testing.T) {
		result := Knit.Swap()

		if result != Purl {
			t.Errorf("Expected Purl, got %v", result)
		}
	})

	t.Run("Returns Knit for Purl", func(t *testing.T) {
		result := Purl.Swap()

		if result != Knit {
			t.Errorf("Expected Knit, got %v", result)
		}
	})
}

func TestParseKnitStitch(t *testing.T) {
	t.Run("returns error for invalid stitch rune", func(t *testing.T) {
		result, err := ParseKnitStitch('🧶')

		checks.CheckHasError(t, result, err, "stitch 🧶 must be a knit (v) or purl (-)")
	})

	t.Run("v results in Knit", func(t *testing.T) {
		result, err := ParseKnitStitch('v')

		checks.CheckHasNoError(t, result, err)
		if result != Knit {
			t.Errorf("Expected Knit, got %v", result)
		}
	})

	t.Run("- results in Purl", func(t *testing.T) {
		result, err := ParseKnitStitch('-')

		checks.CheckHasNoError(t, result, err)
		if result != Purl {
			t.Errorf("Expected Purl, got %v", result)
		}
	})
}
