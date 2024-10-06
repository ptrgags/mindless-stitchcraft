package knitting

import (
	"slices"
)

type Row []KnitStitch
type Fabric []Row

// Take a row of Knits and Purls and swap each stitch type,
// returning a new row.
func (row Row) SwapKnitsAndPurls() Row {
	result := make(Row, len(row))
	for i, stitch := range row {
		result[i] = stitch.Swap()
	}
	return result
}

// Reverse a string.
func (row Row) Reverse() Row {
	result := make(Row, len(row))
	copy(result, row)
	slices.Reverse(result)
	return result
}

func (row Row) ToString() string {
	runes := make([]rune, len(row))
	for i, stitch := range row {
		runes[i] = stitch.ToRune()
	}
	return string(runes)
}

func (fabric Fabric) Rotate180() Fabric {
	n := len(fabric)
	rotated := make(Fabric, n)
	for i, row := range fabric {
		rotated[n-1-i] = row.Reverse()
	}

	return rotated
}

func (fabric Fabric) ToStrings() []string {
	result := make([]string, len(fabric))
	for i, row := range fabric {
		result[i] = row.ToString()
	}
	return result
}
