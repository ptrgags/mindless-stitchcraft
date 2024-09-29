package repeat

import (
	"testing"

	"github.com/ptrgags/mindless-stitchcraft/bracelets"
	"github.com/ptrgags/mindless-stitchcraft/checks"
)

func TestGenerateUncoloredPattern(t *testing.T) {
	t.Run("Zero strands results in error", func(t *testing.T) {
		noStrands := uint(0)

		result, err := GenerateUncoloredPattern(noStrands, []bracelets.Knot{bracelets.BackwardForwardKnot, bracelets.ForwardBackwardKnot})

		checks.CheckHasError(t, result, err, "strandCount must be at least 2")
	})

	t.Run("Odd number of strands returns error", func(t *testing.T) {
		oddStrands := uint(5)

		result, err := GenerateUncoloredPattern(oddStrands, []bracelets.Knot{bracelets.BackwardForwardKnot, bracelets.ForwardBackwardKnot})

		checks.CheckHasError(t, result, err, "strandCount must be an even number")
	})

	t.Run("Valid length and motif does not produce error", func(t *testing.T) {
		length := uint(6)
		motif, _ := bracelets.ParseKnots(`//\\`)

		result, err := GenerateUncoloredPattern(length, motif)

		checks.CheckHasNoError(t, result, err)
	})

	t.Run("Motif that fits in the first pair of rows formats pattern correctly", func(t *testing.T) {
		length := uint(6)
		fitsNicely, _ := bracelets.ParseKnots("///><")

		result, err := GenerateUncoloredPattern(length, fitsNicely)

		expected := []string{
			"/ / /",
			" > < ",
		}
		checks.CheckHasNoError(t, result, err)
		checks.CheckStringGridsEqual(t, result, expected)
	})

	t.Run("Motif that takes multiple rows formats pattern correctly", func(t *testing.T) {
		length := uint(6)
		fitsNicely, _ := bracelets.ParseKnots(`>/\<`)

		result, err := GenerateUncoloredPattern(length, fitsNicely)

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
		checks.CheckHasNoError(t, result, err)
		checks.CheckStringGridsEqual(t, result, expected)
	})
}
