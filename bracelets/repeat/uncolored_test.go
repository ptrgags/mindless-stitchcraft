package repeat

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ptrgags/mindless-stitchcraft/bracelets"
)

func checkRowsEqual(actual []string, expected []string) error {
	if len(actual) != len(expected) {
		return fmt.Errorf("Lengths don't match: len(actual)=%v, len(expected)=%v, actual=%v, expected=%v", len(actual), len(expected), actual, expected)
	}

	for i, actualRow := range actual {
		if actualRow != expected[i] {
			return fmt.Errorf("Row mismatch at position %d: %v, %v actual=%v, expected=%v", i, actualRow, expected[i], actual, expected)
		}
	}

	return nil
}

func TestGenerateUncoloredPattern(t *testing.T) {
	t.Run("Zero strands results in error", func(t *testing.T) {
		noStrands := uint(0)

		result, err := GenerateUncoloredPattern(noStrands, []bracelets.Knot{bracelets.BackwardForwardKnot, bracelets.ForwardBackwardKnot})

		expectedError := "strandCount must be at least 2"
		if err == nil || !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error '%v', got (%v, %v)", expectedError, result, err)
		}
	})

	t.Run("Odd number of strands returns error", func(t *testing.T) {
		oddStrands := uint(5)

		result, err := GenerateUncoloredPattern(oddStrands, []bracelets.Knot{bracelets.BackwardForwardKnot, bracelets.ForwardBackwardKnot})

		expectedError := "strandCount must be an even number"
		if err == nil || !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error '%v', got (%v, %v)", expectedError, result, err)
		}
	})

	t.Run("Valid length and motif does not produce error", func(t *testing.T) {
		length := uint(6)
		motif, _ := bracelets.ParseKnots(`//\\`)

		result, err := GenerateUncoloredPattern(length, motif)

		if err != nil {
			t.Errorf("Expected no error, got (%v, %v)", result, err)
		}
	})

	t.Run("Motif that fits in the first pair of rows formats pattern correctly", func(t *testing.T) {
		length := uint(6)
		fitsNicely, _ := bracelets.ParseKnots("///><")

		result, _ := GenerateUncoloredPattern(length, fitsNicely)

		expected := []string{
			"/ / /",
			" > < ",
		}
		equalsErr := checkRowsEqual(result, expected)
		if equalsErr != nil {
			t.Errorf("%v", equalsErr.Error())
		}
	})

	t.Run("Motif that takes multiple rows formats pattern correctly", func(t *testing.T) {
		length := uint(6)
		fitsNicely, _ := bracelets.ParseKnots(`>/\<`)

		result, _ := GenerateUncoloredPattern(length, fitsNicely)

		expected := []string{
			`> / \`,
			` < > `,
			`/ \ <`,
			` > / `,
			`\ < >`,
			` / \ `,
			`< > /`,
			` \ < `,
		}
		equalsErr := checkRowsEqual(result, expected)
		if equalsErr != nil {
			t.Errorf("%v", equalsErr.Error())
		}
	})
}
