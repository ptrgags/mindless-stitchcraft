package zigzag

import (
	"errors"
	"strings"

	"github.com/ptrgags/mindless-stitchcraft/knitting"
)

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
			result[i] = knitting.SwapKnitsAndPurls(knitting.ReverseRow(row))
		}
	}

	return result
}

func GenerateZigzagPattern(motif string, fabricWidth int) ([]string, error) {
	if err := knitting.ValidateMotif(motif); err != nil {
		return nil, err
	}

	if fabricWidth < 1 {
		return nil, errors.New("fabricWidth must be a positive integer")
	}

	rows := generateRawPattern(motif, fabricWidth)
	rows = ensureEvenRowCount(rows)
	rows = handleReverseRows(rows)
	rows = knitting.Rotate180(rows)

	return rows, nil
}
