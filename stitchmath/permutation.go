package stitchmath

import (
	"errors"
	"fmt"
)

type Permutation struct {
	values []uint
}

func MakeIdentity(length int) Permutation {
	values := make([]uint, length)
	for i := 0; i < length; i++ {
		values[i] = uint(i)
	}

	return Permutation{values}
}

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

func (p Permutation) ElementCount() int {
	return len(p.values)
}

func (p Permutation) GetValues() []uint {
	clone := make([]uint, len(p.values))
	copy(clone, p.values)
	return clone
}

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

func (perm Permutation) Order() uint {
	cycles := perm.CycleDecomposition()

	order := uint(1)
	for _, cycle := range cycles {
		cycleLength := uint(len(cycle))
		order = lcm(order, cycleLength)
	}

	return order
}

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
