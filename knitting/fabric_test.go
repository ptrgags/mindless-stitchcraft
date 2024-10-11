package knitting

import (
	"testing"

	"github.com/ptrgags/mindless-stitchcraft/checks"
)

func TestRowSwapKnitsAndPurls(t *testing.T) {
	t.Run("empty row returns empty row", func(t *testing.T) {
		empty := Row{}

		result := empty.SwapKnitsAndPurls()

		checks.CheckSliceEmpty(t, result)
	})

	t.Run("Swaps knits and purls", func(t *testing.T) {
		row := Row{Knit, Purl, Purl, Knit, Purl}

		result := row.SwapKnitsAndPurls()

		expectedResult := Row{Purl, Knit, Knit, Purl, Knit}
		checks.CheckSlicesEqual(t, result, expectedResult)
	})

	t.Run("Does not modify input row", func(t *testing.T) {
		row := Row{Knit, Purl, Purl}

		row.SwapKnitsAndPurls()

		expectedRow := Row{Knit, Purl, Purl}
		checks.CheckSlicesEqual(t, row, expectedRow)
	})
}

func TestRowReverse(t *testing.T) {
	t.Run("empty row returns empty row", func(t *testing.T) {
		empty := Row{}

		result := empty.Reverse()

		checks.CheckSliceEmpty(t, result)
	})

	t.Run("Reverses order of stitches in input", func(t *testing.T) {
		forward := Row{Purl, Knit, Knit}

		result := forward.Reverse()

		expected := Row{Knit, Knit, Purl}
		checks.CheckSlicesEqual(t, result, expected)
	})

	t.Run("Does not modify input row", func(t *testing.T) {
		row := Row{Knit, Purl, Purl}

		row.Reverse()

		expectedRow := Row{Knit, Purl, Purl}
		checks.CheckSlicesEqual(t, row, expectedRow)
	})
}

func toStitchArray(fabric Fabric) [][]KnitStitch {
	result := make([][]KnitStitch, len(fabric))
	for i, row := range fabric {
		result[i] = []KnitStitch(row)
	}

	return result
}

func TestFabricRotate180(t *testing.T) {
	t.Run("Empty fabric returns empty fabric", func(t *testing.T) {
		empty := Fabric{}

		result := empty.Rotate180()

		checks.CheckSliceEmpty(t, result)
	})

	t.Run("returns stitches rotated 180 degrees", func(t *testing.T) {
		original := Fabric{
			{Knit, Purl, Purl},
			{Knit, Purl, Knit},
			{Purl, Knit, Purl},
		}

		result := original.Rotate180()

		expected := Fabric{
			{Purl, Knit, Purl},
			{Knit, Purl, Knit},
			{Purl, Purl, Knit},
		}
		checks.CheckNestedSlicesEqual(t, toStitchArray(result), toStitchArray(expected))
	})
}

func TestFabricToStrings(t *testing.T) {
	t.Run("Empty fabric returns empty slice", func(t *testing.T) {
		empty := Fabric{}

		result := empty.ToStrings()

		checks.CheckSliceEmpty(t, result)
	})

	t.Run("Stringifies input fabric", func(t *testing.T) {
		fabric := Fabric{
			{Knit, Purl, Purl},
			{Purl, Knit, Purl},
			{Purl, Purl, Knit},
		}

		result := fabric.ToStrings()

		expected := []string{
			"v--",
			"-v-",
			"--v",
		}
		checks.CheckSlicesEqual(t, result, expected)
	})
}
