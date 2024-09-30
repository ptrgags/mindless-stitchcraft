package bracelets

import "fmt"

type Knot int

const (
	ForwardKnot Knot = iota
	BackwardKnot
	ForwardBackwardKnot
	BackwardForwardKnot
)

var knotsToRune = map[Knot]rune{
	ForwardKnot:         '\\',
	BackwardKnot:        '/',
	ForwardBackwardKnot: '>',
	BackwardForwardKnot: '<',
}

var runesToKnots = map[rune]Knot{
	'\\': ForwardKnot,
	'/':  BackwardKnot,
	'>':  ForwardBackwardKnot,
	'<':  BackwardForwardKnot,
}

func (knot Knot) ToRune() (rune, error) {
	if r, ok := knotsToRune[knot]; ok {
		return r, nil
	}

	return rune(0), fmt.Errorf("unknown knot %v", knot)
}

func (knot Knot) SwapsStrands() bool {
	return knot == ForwardKnot || knot == BackwardKnot
}

type VisibleStrand int

const (
	LeftStrand = iota
	RightStrand
)

func (knot Knot) GetVisibleStrand() VisibleStrand {
	if knot == ForwardKnot || knot == ForwardBackwardKnot {
		return LeftStrand
	}

	return RightStrand
}

func fromRune(r rune) (Knot, error) {
	if knot, ok := runesToKnots[r]; ok {
		return knot, nil
	}

	return -1, fmt.Errorf("unknown knot %s", string(r))
}

func ParseKnots(knotString string) ([]Knot, error) {
	result := []Knot{}
	for _, char := range knotString {
		knot, err := fromRune(char)
		if err != nil {
			return []Knot{}, err
		}

		result = append(result, knot)
	}

	return result, nil
}
