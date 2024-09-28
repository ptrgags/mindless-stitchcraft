package repeat

import (
	"testing"

	"github.com/ptrgags/mindless-stitchcraft/bracelets"
	"github.com/ptrgags/mindless-stitchcraft/checks"
)

func TestGenerateColoredPattern(t *testing.T) {
	t.Run("strandCount greater than alphabet length returns error", func(t *testing.T) {
		strands := uint(28)
		anyMotif, _ := bracelets.ParseKnots("///")

		result, err := GenerateColoredPattern(strands, anyMotif)

		checks.CheckHasError(t, result, err, "strandCount must be at most 26")
	})

	t.Run("odd strandCount returns error", func(t *testing.T) {
		strands := uint(5)
		anyMotif, _ := bracelets.ParseKnots("///")

		result, err := GenerateColoredPattern(strands, anyMotif)

		checks.CheckHasError(t, result, err, "strandCount must be an even number")
	})

	t.Run("zero strandCount returns error", func(t *testing.T) {
		strands := uint(0)
		anyMotif, _ := bracelets.ParseKnots("///")

		result, err := GenerateColoredPattern(strands, anyMotif)

		checks.CheckHasError(t, result, err, "strandCount must be at least 2")
	})

	// I was noticing that the spacing on odd rows is doubled for two strands
	t.Run("Two strand pattern that swaps strands does not have extra spacing", func(t *testing.T) {
		strands := uint(2)
		motifThatSwapsStrands, _ := bracelets.ParseKnots("/")

		result, err := GenerateColoredPattern(strands, motifThatSwapsStrands)

		checks.CheckHasNoError(t, result, err)
		// The pattern will have 4 rows + 4 header/footer rows
		expectedWidth := 3
		expectedHeight := 8
		checks.CheckStringGridShape(t, result, expectedWidth, expectedHeight)
	})

	t.Run("Two strand pattern that does not swap strands does not have extra spacing", func(t *testing.T) {
		strands := uint(2)
		motifThatSwapsStrands, _ := bracelets.ParseKnots(">")

		result, err := GenerateColoredPattern(strands, motifThatSwapsStrands)

		checks.CheckHasNoError(t, result, err)
		// The pattern will have 2 rows + 4 header/footer rows
		expectedWidth := 3
		expectedHeight := 6
		checks.CheckStringGridShape(t, result, expectedWidth, expectedHeight)
	})

	t.Errorf("To be continued...")
}
