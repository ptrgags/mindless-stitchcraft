package bracelets

import "github.com/ptrgags/mindless-stitchcraft/stitchmath"

func EvenRowPermutation(knots []Knot) (stitchmath.Permutation, error) {
	// Every knot involves a pair of adjacent strands. Every strand
	// is used for even rows.
	strandCount := 2 * len(knots)

	permutationValues := make([]uint, strandCount)
	for i, knot := range knots {
		leftStrand := 2 * i
		rightStrand := 2*i + 1
		if SwapsStrands(knot) {
			permutationValues[leftStrand] = uint(rightStrand)
			permutationValues[rightStrand] = uint(leftStrand)
		} else {
			permutationValues[leftStrand] = uint(leftStrand)
			permutationValues[rightStrand] = uint(rightStrand)
		}
	}

	return stitchmath.MakePermutation(permutationValues)
}

func OddRowPermutation(knots []Knot) (stitchmath.Permutation, error) {
	// For odd rows, the leftmost and rightmost strands stay in place
	// until the next row, hence the + 2
	strandCount := 2*len(knots) + 2

	permutationValues := make([]uint, strandCount)

	// First and last strands stay in place
	permutationValues[0] = 0
	permutationValues[strandCount-1] = uint(strandCount - 1)

	for i, knot := range knots {
		// the +1 is due to the offset from the fixed strand at position 0
		leftStrand := 2*i + 1
		rightStrand := 2*i + 2
		if SwapsStrands(knot) {
			permutationValues[leftStrand] = uint(rightStrand)
			permutationValues[rightStrand] = uint(leftStrand)
		} else {
			permutationValues[leftStrand] = uint(leftStrand)
			permutationValues[rightStrand] = uint(rightStrand)
		}
	}

	return stitchmath.MakePermutation(permutationValues)
}

func GetPermutations(knotRows [][]Knot) ([]stitchmath.Permutation, error) {
	result := make([]stitchmath.Permutation, len(knotRows))
	for i, row := range knotRows {
		var err error
		if i%2 == 0 {
			result[i], err = EvenRowPermutation(row)
		} else {
			result[i], err = OddRowPermutation(row)
		}

		if err != nil {
			return []stitchmath.Permutation{}, err
		}
	}

	return result, nil
}

/*
func GetColoredPattern(knotRows [][]Knot) ([][]uint, error) {
	if len(knotRows) == 0 {
		return [][]uint{}, nil
	}

	permutations, err := GetPermutations(knotRows)
	if err != nil {
		return [][]uint{}, err
	}

	n := stitchmath.ElementCount(permutations[0])
	current := stitchmath.MakeIdentity(n)

	result := make([][]uint, len(knotRows))
	for i, row := range knotRows {
		knotCount := len(row)
		result[i] = make([]uint, knotCount)
	}

	return result, nil
}
*/
