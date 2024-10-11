package zigzag

import (
	"errors"

	"github.com/ptrgags/mindless-stitchcraft/knitting"
)

// Given a row of the fabric with a fixed width, fill it with copies of motif
// plus a substring that went past the end of the last row by overhang stitches.
//
// This returns (row, overhang) for this row.
func fillRow(overhang int, motif knitting.Motif, fabricWidth int) (knitting.Row, int) {

	// The actual string that went past the end of the last row, it represents
	// the start of this row.
	overhangStr := knitting.Row(motif[len(motif)-overhang:])

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

	left := overhangStr
	middle := knitting.Row(motif.Repeat(uint(repeats)))
	right := knitting.Row(motif[:remaining])

	row := left
	row = append(row, middle...)
	row = append(row, right...)

	return row, nextOverhang
}

// Generate a pattern based on repeating the motif until it aligns
// with the width of the fabric. This works even if the motif is longer
// than a single row!
func generateRawPattern(motif knitting.Motif, fabricWidth int) knitting.Fabric {
	overhang := 0
	fabric := knitting.Fabric{}
	for {
		var row knitting.Row
		row, overhang = fillRow(overhang, motif, fabricWidth)
		fabric = append(fabric, row)
		if overhang == 0 {
			break
		}
	}

	return fabric
}

// Possibly repeat the entire rows twice to ensure the pattern has
// an even number of rows. This ensures that the pattern returns to the
// starting place when knit.
//
// The input rows is treated as an immutable slice, a new slice is always
// allocated.
func ensureEvenRowCount(fabric knitting.Fabric) knitting.Fabric {
	n := len(fabric)
	if n%2 == 0 {
		result := make(knitting.Fabric, n)
		_ = copy(result, fabric)
		return result
	}

	result := make(knitting.Fabric, 0, 2*len(fabric))
	result = append(result, fabric...)
	result = append(result, fabric...)
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
func handleReverseRows(fabric knitting.Fabric) knitting.Fabric {
	result := make(knitting.Fabric, len(fabric))
	_ = copy(result, fabric)

	for i, row := range fabric {
		if i%2 == 0 {
			result[i] = row
		} else {
			result[i] = row.Reverse().SwapKnitsAndPurls()
		}
	}

	return result
}

func GenerateZigzagPattern(motif knitting.Motif, fabricWidth int) ([]string, error) {
	if fabricWidth < 1 {
		return nil, errors.New("fabricWidth must be a positive integer")
	}

	fabric := generateRawPattern(motif, fabricWidth)
	fabric = ensureEvenRowCount(fabric)
	fabric = handleReverseRows(fabric)
	fabric = fabric.Rotate180()

	return fabric.ToStrings(), nil
}
