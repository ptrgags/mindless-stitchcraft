package sync

import (
	"testing"

	"github.com/ptrgags/mindless-stitchcraft/checks"
	"github.com/ptrgags/mindless-stitchcraft/knitting"
)

func TestGeneratePattern(t *testing.T) {
	t.Run("fabricWidth = zero results in error", func(t *testing.T) {
		noFabricWidth := uint(0)
		motif, _ := knitting.ParseMotif("v--")

		result, err := GeneratePattern(noFabricWidth, []knitting.Motif{motif})

		checks.CheckHasError(t, result, err, "fabricWidth must be positive")
	})

	t.Run("No motifs results in error", func(t *testing.T) {
		fabricWidth := uint(10)

		result, err := GeneratePattern(fabricWidth, []knitting.Motif{})

		checks.CheckHasError(t, result, err, "motifs must be non-empty")
	})

	t.Run("short motif is repeated up to fabric width", func(t *testing.T) {
		fabricWidth := uint(10)
		shortMotif, _ := knitting.ParseMotif("v--")

		result, err := GeneratePattern(fabricWidth, []knitting.Motif{shortMotif})

		expectedPattern := []string{
			"-vv-vv-vv-",
			"v--v--v--v",
		}
		checks.CheckHasNoError(t, result, err)
		checks.CheckStringGridsEqual(t, result, expectedPattern)
	})

	t.Run("Long motif is truncated to the fabric width", func(t *testing.T) {
		fabricWidth := uint(4)
		shortMotif, _ := knitting.ParseMotif("--vvv--")

		result, err := GeneratePattern(fabricWidth, []knitting.Motif{shortMotif})

		expectedPattern := []string{
			"vv--",
			"vv--",
		}
		checks.CheckHasNoError(t, result, err)
		checks.CheckStringGridsEqual(t, result, expectedPattern)
	})
}
