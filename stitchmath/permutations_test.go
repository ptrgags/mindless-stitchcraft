package stitchmath

import (
	"fmt"
	"testing"

	"github.com/ptrgags/mindless-stitchcraft/checks"
)

func TestMakePermutation(t *testing.T) {
	t.Run("Zero length permutation returns error", func(t *testing.T) {
		perm, err := MakePermutation([]uint{})

		checks.CheckHasError(t, perm, err, "must have at least one entry")
	})

	t.Run("Out-of-bounds entry results in error", func(t *testing.T) {
		values := []uint{1, 2, 3, 4}

		perm, err := MakePermutation(values)

		checks.CheckHasError(t, perm, err, "values must be in the range [0, 3]")
	})

	t.Run("duplicate entry results in error", func(t *testing.T) {
		values := []uint{0, 0, 1, 2}

		perm, err := MakePermutation(values)

		checks.CheckHasError(t, perm, err, "each entry must be listed exactly once")
	})

	t.Run("Valid permutation does not produce error", func(t *testing.T) {
		values := []uint{0, 2, 1, 3}

		perm, err := MakePermutation(values)

		checks.CheckHasNoError(t, perm, err)
	})
}

func TestElementCount(t *testing.T) {
	t.Run("Computes the number of elements", func(t *testing.T) {
		perm, _ := MakePermutation([]uint{0, 3, 2, 4, 1})

		result := perm.ElementCount()

		expectedLength := 5
		if result != expectedLength {
			t.Errorf("Expected %d, got %v", expectedLength, result)
		}
	})
}

func TestApplyPermutation(t *testing.T) {
	t.Run("out-of-range value is unmodified", func(t *testing.T) {
		perm, _ := MakePermutation([]uint{0, 2, 1, 3})
		var outOfRange uint = 6

		result := perm.Apply(outOfRange)

		if result != outOfRange {
			t.Errorf("Expected %v, got %v", outOfRange, result)
		}
	})

	t.Run("in-range value is cycled correctly", func(t *testing.T) {
		perm, _ := MakePermutation([]uint{0, 2, 1, 3})
		var shouldBeScrambled uint = 1

		result := perm.Apply(shouldBeScrambled)

		var expected uint = 2
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("in-range value is fixed correctly", func(t *testing.T) {
		// This permutation only swaps 1 and 2, so 0 and 3 should remain
		// fixed.
		perm, _ := MakePermutation([]uint{0, 2, 1, 3})
		var shouldBeFixed uint = 3

		result := perm.Apply(shouldBeFixed)

		if result != shouldBeFixed {
			t.Errorf("Expected %v, got %v", shouldBeFixed, result)
		}
	})
}

func TestComposePermutations(t *testing.T) {
	t.Run("mismatched permutation lengths results in error", func(t *testing.T) {
		perm3, _ := MakePermutation([]uint{0, 2, 1})
		perm4, _ := MakePermutation([]uint{1, 3, 0, 2})

		result, err := Compose(perm4, perm3)

		checks.CheckHasError(t, result, err, "permutations must have the same length")
	})

	t.Run("composing valid permutations do not produce error", func(t *testing.T) {
		permA, _ := MakePermutation([]uint{1, 2, 0, 3})
		permB, _ := MakePermutation([]uint{0, 2, 3, 1})

		result, err := Compose(permA, permB)

		checks.CheckHasNoError(t, result, err)
	})

	t.Run("permutations are composed from right to left", func(t *testing.T) {
		permA, _ := MakePermutation([]uint{1, 2, 0, 3}) // (0 1 2)
		permB, _ := MakePermutation([]uint{0, 2, 3, 1}) // (1 2 3)

		result, _ := Compose(permA, permB)

		// permutations are applied from right to left, so we have
		// (0 1 2)(1 2 3) = (0 1)(2 3) which is  [1, 0, 3, 2]
		// if it were computed from left to right, you'd get:
		// (0 2)(1 3) which is not what we want here.
		expectedPerm, _ := MakePermutation([]uint{1, 0, 3, 2})
		if !Equals(result, expectedPerm) {
			t.Errorf("Expected %v, got %v", expectedPerm, result)
		}
	})
}

func TestPermutationEquals(t *testing.T) {
	t.Run("different permutation lengths returns false", func(t *testing.T) {
		perm3, _ := MakePermutation([]uint{1, 2, 0})
		perm4, _ := MakePermutation([]uint{2, 1, 0, 3})

		result := Equals(perm3, perm4)

		if result == true {
			t.Errorf("Expected false, got %v", result)
		}
	})

	t.Run("same permutation returns true", func(t *testing.T) {
		permA, _ := MakePermutation([]uint{1, 2, 0})
		permB, _ := MakePermutation([]uint{1, 2, 0})

		result := Equals(permA, permB)

		if result == false {
			t.Errorf("Expected true, got %v", result)
		}
	})

	t.Run("different permutation returns false", func(t *testing.T) {
		permA, _ := MakePermutation([]uint{1, 2, 0})
		permB, _ := MakePermutation([]uint{1, 0, 2})

		result := Equals(permA, permB)

		if result == true {
			t.Errorf("Expected false, got %v", result)
		}
	})
}

func checkCycleDecompositionsEqual(actualCycles [][]uint, expectedCycles [][]uint) error {
	if len(actualCycles) != len(expectedCycles) {
		return fmt.Errorf("Lengths don't match! actual=%v, expected=%v", actualCycles, expectedCycles)
	}

	for i, actualCycle := range actualCycles {
		expectedCycle := expectedCycles[i]

		if len(actualCycle) != len(expectedCycle) {
			return fmt.Errorf("cycle %d lengths don't match!: %v vs %v, actual=%v, expected=%v", i, actualCycle, expectedCycle, actualCycles, expectedCycles)
		}

		for j, actualElement := range actualCycle {
			if actualElement != expectedCycle[j] {
				return fmt.Errorf("Mismatch at cycle %d, element %d: %v != %v, actual=%v, expected=%v", i, j, actualElement, expectedCycle[j], actualCycles, expectedCycles)
			}
		}
	}

	return nil

}

func TestPermutationCycleDecomposition(t *testing.T) {
	t.Run("Identity permutation returns each element as 1-cycles", func(t *testing.T) {
		identity, _ := MakePermutation([]uint{0, 1, 2, 3, 4})

		result := identity.CycleDecomposition()

		expectedRows := [][]uint{
			{0},
			{1},
			{2},
			{3},
			{4},
		}
		checkErr := checkCycleDecompositionsEqual(result, expectedRows)
		if checkErr != nil {
			t.Error(checkErr.Error())
		}
	})

	t.Run("Involution has one pair and the rest 1-cycles", func(t *testing.T) {
		involution, _ := MakePermutation([]uint{1, 0, 2, 3})

		result := involution.CycleDecomposition()

		expectedCycles := [][]uint{
			{0, 1},
			{2},
			{3},
		}
		checkErr := checkCycleDecompositionsEqual(result, expectedCycles)
		if checkErr != nil {
			t.Error(checkErr.Error())
		}
	})

	t.Run("Cycle through all elements returns a single cycle", func(t *testing.T) {
		cycleBackwards, _ := MakePermutation([]uint{3, 0, 1, 2})

		result := cycleBackwards.CycleDecomposition()

		expectedCycles := [][]uint{
			{0, 3, 2, 1},
		}
		checkErr := checkCycleDecompositionsEqual(result, expectedCycles)
		if checkErr != nil {
			t.Error(checkErr.Error())
		}
	})

	t.Run("Multiple cycles are deinterleaved", func(t *testing.T) {
		multipleCycles, _ := MakePermutation([]uint{2, 5, 4, 1, 0, 3})

		result := multipleCycles.CycleDecomposition()

		expectedCycles := [][]uint{
			{0, 2, 4},
			{1, 5, 3},
		}
		checkErr := checkCycleDecompositionsEqual(result, expectedCycles)
		if checkErr != nil {
			t.Error(checkErr.Error())
		}
	})
}

func TestPermutationOrder(t *testing.T) {
	t.Run("Identity permutation has order 1", func(t *testing.T) {
		// |1| = 1
		identity, _ := MakePermutation([]uint{0, 1, 2, 3, 4})

		order := identity.Order()

		var expectedOrder uint = 1
		if order != expectedOrder {
			t.Errorf("Expected %v, got %v", expectedOrder, order)
		}
	})

	t.Run("Simple swap has order 2", func(t *testing.T) {
		// |(1 2)| = 2
		swap, _ := MakePermutation([]uint{0, 2, 1, 3, 4})

		order := swap.Order()

		var expectedOrder uint = 2
		if order != expectedOrder {
			t.Errorf("Expected %v, got %v", expectedOrder, order)
		}
	})

	t.Run("Simple cycle returns length of cycle", func(t *testing.T) {
		// |(1 2 3)| = 3
		cycle, _ := MakePermutation([]uint{0, 2, 3, 1, 4})

		order := cycle.Order()

		var expectedOrder uint = 3
		if order != expectedOrder {
			t.Errorf("Expected %v, got %v", expectedOrder, order)
		}
	})

	t.Run("Disjoint swaps have order 2", func(t *testing.T) {
		// Two independent swap will both swap back to identity if
		// applied twice.
		// |(0 3)(1 2)| = 2
		twoSwaps, _ := MakePermutation([]uint{3, 2, 1, 0, 4})

		order := twoSwaps.Order()

		var expectedOrder uint = 2
		if order != expectedOrder {
			t.Errorf("Expected %v, got %v", expectedOrder, order)
		}
	})

	t.Run("multiple cycles will have an order based on the lcm of their lengths", func(t *testing.T) {
		// |(0 1)(2 3 4)(5 6 7 8)| = lcm(2, 3, 4) = lcm(6, 4) = 12
		twoSwaps, _ := MakePermutation([]uint{1, 0, 3, 4, 2, 6, 7, 8, 5})

		order := twoSwaps.Order()

		var expectedOrder uint = 12
		if order != expectedOrder {
			t.Errorf("Expected %v, got %v", expectedOrder, order)
		}
	})
}
