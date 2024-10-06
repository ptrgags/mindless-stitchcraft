package knitting

import (
	"errors"
	"unicode/utf8"
)

type Motif []KnitStitch

func (motif Motif) Repeat(n uint) []KnitStitch {
	m := len(motif)
	outputLength := m * int(n)
	result := make([]KnitStitch, outputLength)
	for i := 0; i < outputLength; i++ {
		result[i] = motif[i%m]
	}

	return result
}

func (motif Motif) RepeatToLength(width uint) []KnitStitch {
	result := make([]KnitStitch, width)
	for i := 0; i < int(width); i++ {
		result[i] = motif[i%len(motif)]
	}

	return result
}

func ParseMotif(motif string) (Motif, error) {
	runeCount := utf8.RuneCountInString(motif)

	if runeCount == 0 {
		return Motif{}, errors.New("motif must not be empty")
	}

	result := make(Motif, runeCount)
	for i, r := range motif {
		stitch, err := ParseKnitStitch(r)
		if err != nil {
			return Motif{}, err
		}
		result[i] = stitch
	}

	return result, nil
}
