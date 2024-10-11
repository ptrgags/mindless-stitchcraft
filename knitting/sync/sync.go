package sync

import (
	"errors"

	"github.com/ptrgags/mindless-stitchcraft/knitting"
)

func GeneratePattern(fabricWidth uint, motifs []knitting.Motif) ([]string, error) {
	if fabricWidth < 1 {
		return []string{}, errors.New("fabricWidth must be positive")
	}

	length := len(motifs)
	if length < 1 {
		return []string{}, errors.New("motifs must be non-empty")
	}

	if length%2 == 1 {
		length *= 2
	}

	fabric := make(knitting.Fabric, length)
	for i := 0; i < length; i++ {
		motif := motifs[i%len(motifs)]
		row := knitting.Row(motif.RepeatToLength(fabricWidth))

		if i%2 == 1 {
			row = row.Reverse().SwapKnitsAndPurls()
		}

		fabric[i] = row
	}

	return fabric.Rotate180().ToStrings(), nil
}
