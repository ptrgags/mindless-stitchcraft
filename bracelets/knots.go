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

func ToRune(knot Knot) (rune, error) {
	if r, ok := knotsToRune[knot]; ok {
		return r, nil
	}

	return rune(0), fmt.Errorf("unknown knot %v", knot)
}

func FromRune(r rune) (Knot, error) {
	if knot, ok := runesToKnots[r]; ok {
		return knot, nil
	}

	return -1, fmt.Errorf("unknown knot rune %v", r)
}

func ParseKnots(knotString string) ([]Knot, error) {
	result := []Knot{}
	for _, char := range knotString {
		knot, err := FromRune(char)
		if err != nil {
			return []Knot{}, err
		}

		result = append(result, knot)
	}

	return result, nil
}
