package knitting

import (
	"testing"

	"github.com/ptrgags/mindless-stitchcraft/checks"
)

func TestRepeat(t *testing.T) {
	t.Run("n = 0 returns empty slice", func(t *testing.T) {
		motif, _ := ParseMotif("v--")

		result := motif.Repeat(0)

		checks.CheckSlicesEqual(t, result, []KnitStitch{})
	})

	t.Run("repeat repeats the motif n times", func(t *testing.T) {
		motif, _ := ParseMotif("v--")

		result := motif.Repeat(3)

		expected := []KnitStitch{
			Knit, Purl, Purl,
			Knit, Purl, Purl,
			Knit, Purl, Purl,
		}
		checks.CheckSlicesEqual(t, result, expected)
	})
}

func TestRepeatToLength(t *testing.T) {
	t.Run("zero width returns empty slice", func(t *testing.T) {
		motif, _ := ParseMotif("v--")

		result := motif.RepeatToLength(0)

		checks.CheckSlicesEqual(t, result, []KnitStitch{})
	})

	t.Run("Short width truncates motif", func(t *testing.T) {
		motif, _ := ParseMotif("vvv---")

		result := motif.RepeatToLength(4)

		expected := []KnitStitch{
			Knit, Knit, Knit, Purl,
		}
		checks.CheckSlicesEqual(t, result, expected)
	})

	t.Run("Long width repeats motif", func(t *testing.T) {
		motif, _ := ParseMotif("-v-")

		result := motif.RepeatToLength(7)

		expected := []KnitStitch{
			Purl, Knit, Purl,
			Purl, Knit, Purl,
			Purl,
		}
		checks.CheckSlicesEqual(t, result, expected)
	})
}

func TestParseMotif(t *testing.T) {
	t.Run("empty motif returns error", func(t *testing.T) {
		result, err := ParseMotif("")

		checks.CheckHasError(t, result, err, "motif must not be empty")
	})

	t.Run("Motif with invalid characters returns error", func(t *testing.T) {
		badMotif := "vv--🧶"

		result, err := ParseMotif(badMotif)

		checks.CheckHasError(t, result, err, "stitch 🧶 must be a knit (v) or purl (-)")
	})

	t.Run("Parses valid motif into knits and purls", func(t *testing.T) {
		motif := "v--v"

		result, err := ParseMotif(motif)

		checks.CheckHasNoError(t, result, err)

		expectedMotif := Motif{
			Knit,
			Purl,
			Purl,
			Knit,
		}
		checks.CheckSlicesEqual(t, result, expectedMotif)
	})
}
