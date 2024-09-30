package stitchmath

import (
	"errors"
	"fmt"
)

// A mathematical permutation
// It stores values as if in one-row notation.
// e.g. If the value at index i is v, then this
// means element i is sent to v when applying the permutation.
type Permutation struct {
	values []uint
}

// Make the identity permutation on length elements
func MakeIdentity(length int) Permutation {
	values := make([]uint, length)
	for i := 0; i < length; i++ {
		values[i] = uint(i)
	}

	return Permutation{values}
}

// Make a permutation from the values in one-row notation.
// E.g. the permutation []uint{0, 2, 1, 3} sends
// 0 -> 0
// 1 <-> 2
// 3 -> 3
//
// This constructor takes ownership of values
func MakePermutation(values []uint) (Permutation, error) {
	n := uint(len(values))
	if n == 0 {
		return Permutation{}, errors.New("values must have at least one entry")
	}

	entrySet := make(map[uint]bool)
	for _, x := range values {
		if x >= n {
			return Permutation{}, fmt.Errorf("values must be in the range [0, %v]", n-1)
		}
		entrySet[x] = true
	}

	if uint(len(entrySet)) != n {
		return Permutation{}, fmt.Errorf("each entry must be listed exactly once")
	}

	return Permutation{values}, nil
}

// Get the number of elements n in the permutation.
func (p Permutation) ElementCount() int {
	return len(p.values)
}

// Get a copy of the underlying array, which is stored
// in one-row notation
func (p Permutation) GetValues() []uint {
	clone := make([]uint, len(p.values))
	copy(clone, p.values)
	return clone
}

// Apply the permutation to a single element
func (perm Permutation) Apply(value uint) uint {
	if value > uint(len(perm.values)) {
		return value
	}

	return perm.values[value]
}

func gcd(a uint, b uint) uint {
	if b > a {
		return gcd(b, a)
	}

	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func lcm(a uint, b uint) uint {
	return a * b / gcd(a, b)
}

func findCycle(values []uint, start_index int, visited []bool) []uint {
	n := len(values)
	// The longest possible cycles use N elements.
	result := make([]uint, 1, n)
	result[0] = uint(start_index)

	current_index := start_index
	for i := 0; i < n; i++ {
		element := values[uint(current_index)]
		visited[int(element)] = true

		if element == result[0] {
			break
		}

		result = append(result, element)
		current_index = int(element)
	}

	return result

}

// Compute the cycle decomposition for this permutation,
// though in this implementation, fixed elements are included.
// Cycles are always listed in lexicographical order.
//
// For example, the permutation 2 1 0 4 5 3 (math notation)
// has a cycle decomposition of (0 2)(3 4 5) (math notation)
// This method returns this represented as [][]uint{{0, 2}, {1}, {3, 4, 5}}
func (perm Permutation) CycleDecomposition() [][]uint {
	n := len(perm.values)

	// The length will be at most n for an identity permutation, and otherwise
	// less than this
	result := make([][]uint, 0, n)
	visited := make([]bool, n)

	// Cycles are listed from smallest to largest
	for i := 0; i < n; i++ {
		// We already visited this element in a previous iteration
		if visited[i] {
			continue
		}
		visited[i] = true

		result = append(result, findCycle(perm.values, i, visited))
	}

	return result
}

// Compute the order of the permutation, i.e. how many times to apply the permutation
// to return to identity.
func (perm Permutation) Order() uint {
	cycles := perm.CycleDecomposition()

	order := uint(1)
	for _, cycle := range cycles {
		cycleLength := uint(len(cycle))
		order = lcm(order, cycleLength)
	}

	return order
}

// Compose two permutations. Permutations are applied from right to left
// like functions i.e. compose(a(x), b(x)) = a(b(x))
func Compose(a Permutation, b Permutation) (Permutation, error) {
	n := len(a.values)
	if n != len(b.values) {
		return Permutation{}, errors.New("permutations must have the same length")
	}

	resultValues := make([]uint, n)
	for i := 0; i < n; i++ {
		afterB := b.values[i]
		afterA := a.values[afterB]
		resultValues[i] = afterA
	}

	return Permutation{resultValues}, nil
}

// Check that two permutations are equal
func Equals(a Permutation, b Permutation) bool {
	if len(a.values) != len(b.values) {
		return false
	}

	for i := range a.values {
		if a.values[i] != b.values[i] {
			return false
		}
	}

	return true
}
