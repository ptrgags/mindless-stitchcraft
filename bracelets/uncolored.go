package bracelets

import (
	"errors"
	"fmt"
	"strings"
)

func collectKnots(motif []Knot, cursor uint, count uint) []Knot {
	n := uint(len(motif))
	result := make([]Knot, count)
	for i := uint(0); i < count; i++ {
		result[i] = motif[(cursor+i)%n]
	}

	return result
}

// Format [A, B, C, D] as
// "A B C D"
func formatEvenRow(knots []Knot) string {
	knotStrings := make([]string, len(knots))
	for i, knot := range knots {
		r, _ := ToRune(knot)
		knotStrings[i] = string(r)
	}

	return strings.Join(knotStrings, " ")
}

// Format [A, B, C, D] as
// " A B C D "
func formatOddRow(knots []Knot) string {
	return fmt.Sprintf(" %s ", formatEvenRow(knots))
}

func GenerateUncoloredPattern(strandCount uint, motif []Knot) ([]string, error) {
	if strandCount == 0 {
		return []string{}, errors.New("strandCount must be at least 2")
	}

	if strandCount%2 != 0 {
		return []string{}, errors.New("strandCount must be an even number")
	}

	// Stitches are staggered like this:
	// x x x x
	//  x x x
	evenStitchCount := strandCount / 2
	oddStitchCount := evenStitchCount - 1

	result := []string{}
	cursor := uint(0)
	for i := 0; i < len(motif); i++ {
		evenKnots := collectKnots(motif, cursor, evenStitchCount)
		evenRow := formatEvenRow(evenKnots)
		cursor += evenStitchCount

		oddKnots := collectKnots(motif, cursor, oddStitchCount)
		oddRow := formatOddRow(oddKnots)
		cursor += oddStitchCount

		cursor %= uint(len(motif))

		// Detect pattern repeat
		if i > 0 && cursor == 0 {
			break
		}

		result = append(result, evenRow, oddRow)

	}

	return result, nil
}
