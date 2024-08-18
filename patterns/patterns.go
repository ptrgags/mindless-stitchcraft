package patterns

import (
	"errors"
	"slices"
	"strings"
)

// Check that a motif is nonempty and is composed of 'v' and '-' runes.
func validateMotif(motif string) error {
	if len(motif) == 0 {
		return errors.New("motif must not be empty")
	}

	for _, r := range motif {
		if r != 'v' && r != '-' {
			return errors.New("motif has invalid characters. It must be a string of knits ('v') and purls ('-')")
		}
	}

	return nil
}

// Given a row of the fabric with a fixed width, fill it with copies of motif
// plus a substring that went past the end of the last row by overhang stitches.
//
// This returns (row, overhang) for this row.
func fillRow(overhang int, motif string, fabricWidth int) (string, int) {

	// The actual string that went past the end of the last row, it represents
	// the start of this row.
	overhangStr := motif[len(motif)-overhang:]

	// The overhang from the previous row was so long it fills the entire row
	if overhang > fabricWidth {
		return overhangStr[:fabricWidth], overhang - fabricWidth
	}

	repeats := (fabricWidth - overhang) / len(motif)
	remaining := (fabricWidth - overhang) % len(motif)

	nextOverhang := 0
	if remaining != 0 {
		nextOverhang = len(motif) - remaining
	}

	row := overhangStr + strings.Repeat(motif, repeats) + motif[:remaining]
	return row, nextOverhang
}

// Generate a pattern based on repeating the motif until it aligns
// with the width of the fabric. This works even if the motif is longer
// than a single row!
//
// motif must be a valid knitting motif of knits (v) and purls (-)
// fabricWidth must be non-zero
func generateRawPattern(motif string, fabricWidth int) []string {
	overhang := 0
	rows := []string{}
	for {
		var row string
		row, overhang = fillRow(overhang, motif, fabricWidth)
		rows = append(rows, row)
		if overhang == 0 {
			break
		}
	}

	return rows
}

// Possibly repeat the entire rows twice to ensure the pattern has
// an even number of rows. This ensures that the pattern returns to the
// starting place when knit.
//
// The input rows is treated as an immutable slice, a new slice is always
// allocated.
func ensureEvenRowCount(rows []string) []string {
	n := len(rows)
	if n%2 == 0 {
		result := make([]string, n)
		_ = copy(result, rows)
		return result
	}

	result := make([]string, 0, 2*len(rows))
	result = append(result, rows...)
	result = append(result, rows...)
	return result
}

// Take a string of stiches (either 'v' for knit or '-' for purl)
// and swap the knits and purls. This is one part of what happens when
// you flip the fabric over when knitting. This returns a new string
func swapKnitsAndPurls(stitches string) string {
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
func reverse(s string) string {
	result := []rune(s)
	slices.Reverse(result)
	return string(result)
}

// pass in rows of the fabric listed as knit ('v') and purl ('-') as the
// knitter
//
// However, every second row is knit on the wrong side of the fabric, so
// the direction of stitching is reversed, and also knits show up as purls
// and vice-versa on the front. This function takes care of this.
//
// Rows is treated as an immutable slice, so a new slice is allocated.
func handleReverseRows(rows []string) []string {
	result := make([]string, len(rows))
	_ = copy(result, rows)

	for i, row := range rows {
		if i%2 == 0 {
			result[i] = row
		} else {
			result[i] = swapKnitsAndPurls(reverse(row))
		}
	}

	return result
}

func GeneratePattern(motif string, fabricWidth int) ([]string, error) {
	if err := validateMotif(motif); err != nil {
		return nil, err
	}

	if fabricWidth < 1 {
		return nil, errors.New("fabricWidth must be a positive integer")
	}

	rows := generateRawPattern(motif, fabricWidth)
	rows = ensureEvenRowCount(rows)
	rows = handleReverseRows(rows)

	return rows, nil
}

func rotate180(rows []string) []string {
	n := len(rows)
	rotated := make([]string, n)
	for i, row := range rows {
		rotated[n-1-i] = reverse(row)
	}

	return rotated
}

func GeneratePrintablePattern(motif string, fabricWidth int) ([]string, error) {
	result, err := GeneratePattern(motif, fabricWidth)
	if err != nil {
		return nil, err
	}
	return rotate180(result), nil
}
