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
