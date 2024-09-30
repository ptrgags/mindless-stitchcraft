package sync

import "github.com/ptrgags/mindless-stitchcraft/knitting"

func repeatToLength(motif string, fabricWidth int) string {
	runes := []rune(motif)

	row := make([]rune, fabricWidth)
	for i := 0; i < fabricWidth; i++ {
		row[i] = runes[i%len(runes)]
	}

	return string(row)
}

func GeneratePattern(fabricWidth int, motifs []string) ([]string, error) {
	for _, motif := range motifs {
		err := knitting.ValidateMotif(motif)
		if err != nil {
			return []string{}, err
		}
	}

	length := len(motifs)
	if length%2 == 1 {
		length *= 2
	}

	rows := make([]string, length)
	for i := 0; i < length; i++ {
		motif := motifs[i%len(motifs)]
		row := repeatToLength(motif, fabricWidth)

		if i%2 == 1 {
			row = knitting.SwapKnitsAndPurls(knitting.ReverseRow(row))
		}

		rows[i] = row
	}

	rows = knitting.Rotate180(rows)

	return rows, nil
}
