package repeat

import (
	"testing"

	"github.com/ptrgags/mindless-stitchcraft/bracelets"
	"github.com/ptrgags/mindless-stitchcraft/checks"
)

func TestGenerateColoredPattern(t *testing.T) {
	t.Run("odd strandCount returns error", func(t *testing.T) {
		strands := []rune("ABC")
		anyMotif, _ := bracelets.ParseKnots("///")

		result, err := GenerateColoredPattern(strands, anyMotif)

		checks.CheckHasError(t, result, err, "strandCount must be an even number")
	})

	t.Run("zero strandCount returns error", func(t *testing.T) {
		strands := []rune{}
		anyMotif, _ := bracelets.ParseKnots("///")

		result, err := GenerateColoredPattern(strands, anyMotif)

		checks.CheckHasError(t, result, err, "strandCount must be at least 2")
	})

	// I was noticing that the spacing on odd rows is doubled for two strands
	t.Run("Two strand pattern that swaps strands does not have extra spacing", func(t *testing.T) {
		strands := []rune("AB")
		motifThatSwapsStrands, _ := bracelets.ParseKnots("/")

		result, err := GenerateColoredPattern(strands, motifThatSwapsStrands)

		checks.CheckHasNoError(t, result, err)
		// The pattern will have 4 rows + 4 header/footer rows
		expectedWidth := 3
		expectedHeight := 8
		checks.CheckStringGridShape(t, result, expectedWidth, expectedHeight)
	})

	t.Run("Two strand pattern that does not swap strands does not have extra spacing", func(t *testing.T) {
		strands := []rune("AB")
		motifThatSwapsStrands, _ := bracelets.ParseKnots(">")

		result, err := GenerateColoredPattern(strands, motifThatSwapsStrands)

		checks.CheckHasNoError(t, result, err)
		// The pattern will have 2 rows + 4 header/footer rows
		expectedWidth := 3
		expectedHeight := 6
		checks.CheckStringGridShape(t, result, expectedWidth, expectedHeight)
	})

	t.Run("Two strand pattern that swaps strands produces the correct pattern", func(t *testing.T) {
		strands := []rune("AB")
		motifThatSwapsStrands, _ := bracelets.ParseKnots("/")

		result, err := GenerateColoredPattern(strands, motifThatSwapsStrands)

		expectedPattern := []string{
			"A B",
			"| |",
			" B ",
			"B A",
			" A ",
			"A B",
			"| |",
			"A B",
		}
		checks.CheckHasNoError(t, result, err)
		checks.CheckStringGridsEqual(t, result, expectedPattern)
	})

	t.Run("Two strand pattern that does not swap strands produces the correct pattern", func(t *testing.T) {
		strands := []rune("AB")
		motifThatSwapsStrands, _ := bracelets.ParseKnots(">")

		result, err := GenerateColoredPattern(strands, motifThatSwapsStrands)

		expectedPattern := []string{
			"A B",
			"| |",
			" A ",
			"A B",
			"| |",
			"A B",
		}
		checks.CheckHasNoError(t, result, err)
		checks.CheckStringGridsEqual(t, result, expectedPattern)
	})

	t.Run("pattern without repeats produces the correct shape", func(t *testing.T) {
		strands := []rune("ABCD")
		exactFitMotif, _ := bracelets.ParseKnots("<><") // seems fishy 🤔

		result, err := GenerateColoredPattern(strands, exactFitMotif)

		// The motif exactly fits 2 rows, and doesn't swap strands
		expectedWidth := 7
		expectedHeight := 6 // 2 rows + 4 header/footer
		checks.CheckHasNoError(t, result, err)
		checks.CheckStringGridShape(t, result, expectedWidth, expectedHeight)
	})

	t.Run("pattern without repeats produces the correct pattern", func(t *testing.T) {
		strands := []rune("ABCD")
		exactFitMotif, _ := bracelets.ParseKnots("<><") // seems fishy 🤔

		result, err := GenerateColoredPattern(strands, exactFitMotif)

		// The motif exactly fits 2 rows, and doesn't swap strands
		expectedPattern := []string{
			"A B C D",
			"| | | |",
			" B   C ",
			"A  C  D",
			"| | | |",
			"A B C D",
		}
		checks.CheckHasNoError(t, result, err)
		checks.CheckStringGridsEqual(t, result, expectedPattern)
	})

	t.Run("pattern with repeats produces the correct shape", func(t *testing.T) {
		strands := []rune("ABCD")
		exactFitMotif, _ := bracelets.ParseKnots(`//\`)

		result, err := GenerateColoredPattern(strands, exactFitMotif)

		// The motif exactly fits 2 rows, but due to swapping it takes 4 repeats
		// to permute the strands back
		expectedWidth := 7
		expectedHeight := 12 // 4 * 2 rows + 4 header/footer
		checks.CheckHasNoError(t, result, err)
		checks.CheckStringGridShape(t, result, expectedWidth, expectedHeight)
	})

	t.Run("pattern with repeats produces the correct pattern", func(t *testing.T) {
		strands := []rune("ABCD")
		exactFitMotif, _ := bracelets.ParseKnots(`//\`)

		result, err := GenerateColoredPattern(strands, exactFitMotif)

		// The motif exactly fits 2 rows, but due to swapping it takes 4 repeats
		// to permute the strands back
		expectedPattern := []string{
			"A B C D",
			"| | | |",
			" B   D ", // BADC after this row
			"B  A  C", // BDAC
			" D   C ", // DBCA
			"D  B  A", // DCBA
			" C   A ", // CDAB
			"C  D  B", // CADB
			" A   B ", // ACBD
			"A  C  D", // ABCD
			"| | | |",
			"A B C D",
		}
		checks.CheckHasNoError(t, result, err)
		checks.CheckStringGridsEqual(t, result, expectedPattern)
	})

	t.Run("motif longer than 2 rows produces the correct pattern", func(t *testing.T) {
		strands := []rune("ABCD")
		motif, _ := bracelets.ParseKnots(`///\`)

		result, err := GenerateColoredPattern(strands, motif)
		expectedPattern := []string{
			"A B C D",
			"| | | |",
			" B   D ", //  / /  BADC
			"B  D  C", //   /   BDAC
			" B   C ", //  \ /  DBCA
			"D  C  A", //   /   DCBA
			" C   B ", //  / \  CDAB
			"C  A  B", //   /   CADB
			" A   B ", //  / /  ACBD
			"A  C  D", //   \   ABCD
			"| | | |",
			"A B C D",
		}
		checks.CheckHasNoError(t, result, err)
		checks.CheckStringGridsEqual(t, result, expectedPattern)
	})

	t.Run("motif with mixed knots types produces the correct pattern", func(t *testing.T) {
		strands := []rune("ABCD")
		motif, _ := bracelets.ParseKnots(`<>/\`)

		result, err := GenerateColoredPattern(strands, motif)
		expectedPattern := []string{
			"A B C D",
			"| | | |",
			" B   C ", //  < >  ABCD
			"A  C  D", //   /   ACBD
			" A   D ", //  \ <  CABD
			"C  A  D", //   >   CABD
			" A   B ", //  / \  ACDB
			"A  D  B", //   <   ACDB
			" A   B ", //  > /  ACBD
			"A  C  D", //   \   ABCD
			"| | | |",
			"A B C D",
		}
		checks.CheckHasNoError(t, result, err)
		checks.CheckStringGridsEqual(t, result, expectedPattern)
	})

	// Design choice. Easier to implement this way.
	t.Run("Repeats pattern even if strand labels make it redundant", func(t *testing.T) {
		strands := []rune("AABB")
		motif, _ := bracelets.ParseKnots(`//<`)

		result, err := GenerateColoredPattern(strands, motif)
		expectedPattern := []string{
			"A A B B",
			"| | | |", //  to show what's going on with 4 unique strands:
			" A   B ", //  / /  BADC
			"A  B  B", //   <   BADC
			" A   B ", //  / /  ABCD
			"A  B  B", //   <   ABCD
			"| | | |",
			"A A B B",
		}
		checks.CheckHasNoError(t, result, err)
		checks.CheckStringGridsEqual(t, result, expectedPattern)
	})

	t.Run("Pattern with only F/B and B/F knots produces a short pattern", func(t *testing.T) {
		strands := []rune("ABCDEF")
		motif, _ := bracelets.ParseKnots(`>>><<`)

		result, err := GenerateColoredPattern(strands, motif)
		expectedPattern := []string{
			"A B C D E F",
			"| | | | | |",
			" A   C   E ", //  > > >  ABCDEF
			"A  C   E  F", //   < <   ABCDEF
			"| | | | | |",
			"A B C D E F",
		}
		checks.CheckHasNoError(t, result, err)
		checks.CheckStringGridsEqual(t, result, expectedPattern)
	})
}
