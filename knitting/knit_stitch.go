package knitting

import "fmt"

type KnitStitch int

const (
	Knit KnitStitch = iota
	Purl
)

func (stitch KnitStitch) ToRune() rune {
	if stitch == Knit {
		return 'v'
	}

	return '-'
}

func (stitch KnitStitch) Swap() KnitStitch {
	if stitch == Knit {
		return Purl
	}

	return Knit
}

func ParseKnitStitch(stitch rune) (KnitStitch, error) {
	if stitch == 'v' {
		return Knit, nil
	}

	if stitch == '-' {
		return Purl, nil
	}

	return Knit, fmt.Errorf("stitch %s must be a knit (v) or purl (-)", string(stitch))
}
