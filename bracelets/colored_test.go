package bracelets

import (
	"testing"

	"github.com/ptrgags/mindless-stitchcraft/stitchmath"
)

func TestEvenRowPermutation(t *testing.T) {

	t.Run("Empty knots results in error", func(t *testing.T) {
		knots := []Knot{}

		result, err := EvenRowPermutation(knots)

		if err == nil {
			t.Errorf("Expected error, got (%v, %v)", result, err)
		}
	})

	t.Run("Row with knots does not return error", func(t *testing.T) {
		knots := []Knot{ForwardBackwardKnot, BackwardForwardKnot}

		result, err := EvenRowPermutation(knots)

		if err != nil {
			t.Errorf("Expected no error, got (%v, %v)", result, err)
		}
	})

	t.Run("Row with only forward-backward or backward-forward knots returns identity", func(t *testing.T) {
		knots := []Knot{ForwardBackwardKnot, ForwardBackwardKnot, BackwardForwardKnot}

		result, _ := EvenRowPermutation(knots)

		identity := stitchmath.MakeIdentity(6)
		if !stitchmath.Equals(result, identity) {
			t.Errorf("Expected result to be the identity permutation, got %v", result)
		}
	})

	t.Run("Row with only forward or backward knots swaps all strands", func(t *testing.T) {
		knots := []Knot{BackwardKnot, ForwardKnot, BackwardKnot, ForwardKnot}

		result, _ := EvenRowPermutation(knots)

		// Swap each pair of adjacent strands
		identity, _ := stitchmath.MakePermutation([]uint{
			1, 0, 3, 2, 5, 4, 7, 6,
		})

		if !stitchmath.Equals(result, identity) {
			t.Errorf("Expected result to be the identity permutation, got %v", result)
		}
	})

	t.Run("Row with mixed knots computes the correct permutation", func(t *testing.T) {
		knots := []Knot{ForwardKnot, ForwardBackwardKnot, BackwardKnot, BackwardForwardKnot}

		result, _ := EvenRowPermutation(knots)

		identity, _ := stitchmath.MakePermutation([]uint{
			1, 0, 2, 3, 5, 4, 6, 7,
		})

		if !stitchmath.Equals(result, identity) {
			t.Errorf("Expected result to be the identity permutation, got %v", result)
		}
	})
}

func TestOddRowPermutation(t *testing.T) {
	t.Run("Empty knots results in identity", func(t *testing.T) {
		knots := []Knot{}

		result, err := OddRowPermutation(knots)

		// The left and right strands are unused.
		identity := stitchmath.MakeIdentity(2)
		if !stitchmath.Equals(result, identity) {
			t.Errorf("expected result to be the identity permutation, got (%v, %v)", result, err)
		}
	})

	t.Run("Row with knots does not return error", func(t *testing.T) {
		knots := []Knot{ForwardBackwardKnot, BackwardForwardKnot}

		result, err := OddRowPermutation(knots)

		if err != nil {
			t.Errorf("Expected no error, got (%v, %v)", result, err)
		}
	})

	t.Run("Row with only forward-backward or backward-forward knots returns identity", func(t *testing.T) {
		knots := []Knot{ForwardBackwardKnot, ForwardBackwardKnot, BackwardForwardKnot}

		result, err := OddRowPermutation(knots)

		identity := stitchmath.MakeIdentity(8)
		if !stitchmath.Equals(result, identity) {
			t.Errorf("Expected result to be the identity permutation, got %v, %v", result, err)
		}
	})

	t.Run("Row with only forward or backward knots swaps all strands except first and last", func(t *testing.T) {
		knots := []Knot{BackwardKnot, ForwardKnot, BackwardKnot, ForwardKnot}

		result, err := OddRowPermutation(knots)

		// Swap each pair of adjacent strands
		expected, _ := stitchmath.MakePermutation([]uint{
			0, 2, 1, 4, 3, 6, 5, 8, 7, 9,
		})
		if !stitchmath.Equals(result, expected) {
			t.Errorf("Expected result to be %v, got (%v, %v)", expected, result, err)
		}
	})

	t.Run("Row with mixed knots computes the correct permutation", func(t *testing.T) {
		knots := []Knot{ForwardKnot, ForwardBackwardKnot, BackwardKnot, BackwardForwardKnot}

		result, err := OddRowPermutation(knots)

		expected, _ := stitchmath.MakePermutation([]uint{
			0, 2, 1, 3, 4, 6, 5, 7, 8, 9,
		})
		if !stitchmath.Equals(result, expected) {
			t.Errorf("Expected result to be %v, got (%v, %v)", expected, result, err)
		}
	})
}
