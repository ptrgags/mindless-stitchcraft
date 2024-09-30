package knitting

import (
	"slices"
	"strings"
)

// Take a string of stiches (either 'v' for knit or '-' for purl)
// and swap the knits and purls. This is one part of what happens when
// you flip the fabric over when knitting. This returns a new string
func SwapKnitsAndPurls(stitches string) string {
	return strings.Map(func(r rune) rune {
		if r == 'v' {
			return '-'
		} else if r == '-' {
			return 'v'
		}

		return r
	}, stitches)
}

// Reverse a string.
func ReverseRow(s string) string {
	result := []rune(s)
	slices.Reverse(result)
	return string(result)
}

func Rotate180(rows []string) []string {
	n := len(rows)
	rotated := make([]string, n)
	for i, row := range rows {
		rotated[n-1-i] = ReverseRow(row)
	}

	return rotated
}
