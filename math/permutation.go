package math

import (
	"errors"
	"fmt"
)

type Permutation struct {
	values []uint
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

func Apply(perm Permutation, value uint) uint {
	if value > uint(len(perm.values)) {
		return value
	}

	return perm.values[value]
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
