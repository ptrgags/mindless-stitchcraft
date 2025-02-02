package repeat

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ptrgags/mindless-stitchcraft/bracelets"
	"github.com/ptrgags/mindless-stitchcraft/stitchmath"
)

func evenRowPermutation(knots []bracelets.Knot) (stitchmath.Permutation, error) {
	// Every knot involves a pair of adjacent strands. Every strand
	// is used for even rows.
	strandCount := 2 * len(knots)

	permutationValues := make([]uint, strandCount)
	for i, knot := range knots {
		leftStrand := 2 * i
		rightStrand := 2*i + 1
		if knot.SwapsStrands() {
			permutationValues[leftStrand] = uint(rightStrand)
			permutationValues[rightStrand] = uint(leftStrand)
		} else {
			permutationValues[leftStrand] = uint(leftStrand)
			permutationValues[rightStrand] = uint(rightStrand)
		}
	}

	return stitchmath.MakePermutation(permutationValues)
}

func oddRowPermutation(knots []bracelets.Knot) (stitchmath.Permutation, error) {
	// For odd rows, the leftmost and rightmost strands stay in place
	// until the next row, hence the + 2
	strandCount := 2*len(knots) + 2

	permutationValues := make([]uint, strandCount)

	// First and last strands stay in place
	permutationValues[0] = 0

	for i, knot := range knots {
		// the +1 is due to the offset from the fixed strand at position 0
		leftStrand := 2*i + 1
		rightStrand := 2*i + 2
		if knot.SwapsStrands() {
			permutationValues[leftStrand] = uint(rightStrand)
			permutationValues[rightStrand] = uint(leftStrand)
		} else {
			permutationValues[leftStrand] = uint(leftStrand)
			permutationValues[rightStrand] = uint(rightStrand)
		}
	}

	permutationValues[strandCount-1] = uint(strandCount - 1)

	return stitchmath.MakePermutation(permutationValues)
}

func getPermutations(knotRows [][]bracelets.Knot) ([]stitchmath.Permutation, error) {
	result := make([]stitchmath.Permutation, len(knotRows))
	for i, row := range knotRows {
		var err error
		if i%2 == 0 {
			result[i], err = evenRowPermutation(row)
		} else {
			result[i], err = oddRowPermutation(row)
		}

		if err != nil {
			return []stitchmath.Permutation{}, err
		}
	}

	return result, nil
}

func colorEvenRow(strands []uint, knots []bracelets.Knot) []uint {
	result := make([]uint, len(knots))
	for i, knot := range knots {
		leftStrand := 2 * i
		rightStrand := 2*i + 1

		if knot.GetVisibleStrand() == bracelets.LeftStrand {
			result[i] = strands[leftStrand]
		} else {
			result[i] = strands[rightStrand]
		}
	}
	return result
}

func colorOddRow(strands []uint, knots []bracelets.Knot) []uint {
	// Include the end strands
	n := len(knots) + 2
	result := make([]uint, len(knots)+2)
	// Outermost strands are visible and never swap on odd rows.
	result[0] = strands[0]
	result[n-1] = strands[len(strands)-1]

	for i, knot := range knots {
		leftStrand := 2*i + 1
		rightStrand := 2*i + 2

		if knot.GetVisibleStrand() == bracelets.LeftStrand {
			result[i+1] = strands[leftStrand]
		} else {
			result[i+1] = strands[rightStrand]
		}
	}
	return result
}

func composeAll(perms []stitchmath.Permutation) (stitchmath.Permutation, error) {
	product := perms[0]
	var err error
	for i := 1; i < len(perms); i++ {
		product, err = stitchmath.Compose(perms[i], product)
	}

	return product, err
}

func getColoredPattern(knotRows [][]bracelets.Knot) ([][]uint, error) {
	inputRows := len(knotRows)
	if inputRows == 0 {
		return [][]uint{}, nil
	}

	if inputRows%2 == 1 {
		return [][]uint{}, fmt.Errorf("knotRows must have an even number of rows, got %d", inputRows)
	}

	permutations, err := getPermutations(knotRows)
	if err != nil {
		return [][]uint{}, err
	}

	product, err := composeAll(permutations)
	if err != nil {
		return [][]uint{}, err
	}

	patternRepeats := product.Order()
	resultRowCount := int(patternRepeats) * inputRows

	n := permutations[0].ElementCount()

	// Inverse of the current chain of permutations.
	// The forward permutation computes where each strand color
	// ends up. We want the opposite - for a given strand, which
	// color ended up here? So use the inverse to compute these
	// color labels.
	inversePermutation := stitchmath.MakeIdentity(n)
	result := make([][]uint, resultRowCount)
	for i := 0; i < resultRowCount; i++ {
		strandOrder := inversePermutation.GetValues()
		row := knotRows[i%inputRows]
		permutation := permutations[i%inputRows]

		if i%2 == 0 {
			result[i] = colorEvenRow(strandOrder, row)
		} else {
			result[i] = colorOddRow(strandOrder, row)
		}

		// IMPORTANT - the permutations used here are always involutions,
		// so A^(-1) = A, B^(-1) = B
		// so (AB)^-1 = B^(-1)A^(-1) = BA
		// So reversing the multiplication order computes the inverse product without
		// having to explicitly compute inverses!
		inversePermutation, err = stitchmath.Compose(inversePermutation, permutation)
		if err != nil {
			return [][]uint{}, err
		}
	}

	return result, nil
}

func labelStrands(strandLabels []rune, unlabeledPattern [][]uint) ([][]rune, error) {
	result := make([][]rune, len(unlabeledPattern))
	for i, unlabeledRow := range unlabeledPattern {
		runes := make([]rune, len(unlabeledRow))
		for j, strandIndex := range unlabeledRow {
			runes[j] = strandLabels[int(strandIndex)]
		}
		result[i] = runes
	}

	return result, nil
}

func joinRunes(runes []rune, sep string) string {
	values := make([]string, len(runes))
	for i, value := range runes {
		values[i] = string(value)
	}

	return strings.Join(values, sep)
}

func formatRows(strandLabels []rune, labeledRows [][]rune) []string {
	strandString := joinRunes(strandLabels, " ")

	straightRow := make([]string, len(strandLabels))
	for i := 0; i < len(strandLabels); i++ {
		straightRow[i] = "|"
	}
	straightString := strings.Join(straightRow, " ")

	result := make([]string, len(labeledRows)+4)
	result[0] = strandString
	result[1] = straightString
	for i, row := range labeledRows {
		if i%2 == 0 {
			middle := joinRunes(row, "   ")
			result[2+i] = fmt.Sprintf(" %s ", middle)
		} else {
			// For odd rows, the outer strands have smaller spacing
			first := string(row[0])
			middle := joinRunes(row[1:len(row)-1], "   ")
			last := string(row[len(row)-1])

			if middle == "" {
				result[2+i] = fmt.Sprintf("%s %s", first, last)
			} else {
				result[2+i] = fmt.Sprintf("%s  %s  %s", first, middle, last)
			}
		}
	}
	result[len(result)-2] = straightString
	result[len(result)-1] = strandString

	return result
}

// Create a preview of the friendship bracelet colored with the
// strandLabels. This will repeat the pattern until the strands at the
// bottom equal the strands at the top.
//
// strandLabels is a slice of runes representing the colors. In practice
// usually you have pairs of the same color, so e.g. []rune("ABBA") is valid
//
// motif is the list of knots to repeat. See bracelets.ParseKnots
func GenerateColoredPattern(strandLabels []rune, motif []bracelets.Knot) ([]string, error) {
	strandCount := uint(len(strandLabels))

	if len(motif) == 0 {
		return []string{}, errors.New("motif must have at least one knot")
	}

	knotRows, err := GenerateUncoloredKnots(strandCount, motif)
	if err != nil {
		return []string{}, err
	}

	unlabeledRows, err := getColoredPattern(knotRows)
	if err != nil {
		return []string{}, err
	}

	labeledRows, err := labelStrands(strandLabels, unlabeledRows)
	if err != nil {
		return []string{}, err
	}

	formatted := formatRows(strandLabels, labeledRows)

	return formatted, nil
}
