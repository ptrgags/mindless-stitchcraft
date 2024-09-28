package repeat

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ptrgags/mindless-stitchcraft/bracelets"
)

func collectKnots(motif []bracelets.Knot, cursor uint, count uint) []bracelets.Knot {
	n := uint(len(motif))
	result := make([]bracelets.Knot, count)
	for i := uint(0); i < count; i++ {
		result[i] = motif[(cursor+i)%n]
	}

	return result
}

// Format [A, B, C, D] as
// "A B C D"
func formatEvenRow(knots []bracelets.Knot) string {
	knotStrings := make([]string, len(knots))
	for i, knot := range knots {
		r, _ := knot.ToRune()
		knotStrings[i] = string(r)
	}

	return strings.Join(knotStrings, " ")
}

// Format [A, B, C, D] as
// " A B C D "
func formatOddRow(knots []bracelets.Knot) string {
	return fmt.Sprintf(" %s ", formatEvenRow(knots))
}

func GenerateUncoloredKnots(strandCount uint, motif []bracelets.Knot) ([][]bracelets.Knot, error) {
	if strandCount == 0 {
		return [][]bracelets.Knot{}, errors.New("strandCount must be at least 2")
	}

	if strandCount%2 != 0 {
		return [][]bracelets.Knot{}, errors.New("strandCount must be an even number")
	}

	// Stitches are staggered like this:
	// x x x x
	//  x x x
	evenStitchCount := strandCount / 2
	oddStitchCount := evenStitchCount - 1

	result := [][]bracelets.Knot{}
	// The cursor loops over the motif
	cursor := uint(0)
	for i := 0; i < len(motif); i++ {
		// Detect pattern repeat
		if i > 0 && cursor == 0 {
			break
		}

		evenKnots := collectKnots(motif, cursor, evenStitchCount)
		cursor += evenStitchCount

		oddKnots := collectKnots(motif, cursor, oddStitchCount)
		cursor += oddStitchCount

		cursor %= uint(len(motif))
		result = append(result, evenKnots, oddKnots)
	}

	return result, nil
}

func GenerateUncoloredPattern(strandCount uint, motif []bracelets.Knot) ([]string, error) {
	knotRows, err := GenerateUncoloredKnots(strandCount, motif)
	if err != nil {
		return []string{}, err
	}

	result := []string{}
	for i, knots := range knotRows {
		var formattedRow string
		if i%2 == 0 {
			formattedRow = formatEvenRow(knots)
		} else {
			formattedRow = formatOddRow(knots)
		}
		result = append(result, formattedRow)
	}
	return result, nil
}
