package math

import (
	"errors"
	"fmt"
)

type Permutation struct {
	values []uint
}

func (p Permutation) getLength() int {
	return len(p.values)
}

func makePermutation(values []uint) (Permutation, error) {
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

func apply(perm Permutation, value uint) uint {
	return 0
}

func compose(a Permutation, b Permutation) (Permutation, error) {
	return Permutation{}, nil
}
