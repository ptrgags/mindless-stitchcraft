package bracelets

import (
	"testing"

	"github.com/ptrgags/mindless-stitchcraft/checks"
)

func TestToRune(t *testing.T) {
	t.Run("Invalid knot returns error", func(t *testing.T) {
		knot := ForwardKnot
		knot += 10

		result, err := knot.ToRune()

		checks.CheckHasError(t, result, err, "unknown knot 10")
	})

	t.Run(`Forward knot returns \`, func(t *testing.T) {
		result, err := ForwardKnot.ToRune()

		checks.CheckHasNoError(t, result, err)
		if result != '\\' {
			t.Errorf(`Expected \, got %v`, string(result))
		}
	})

	t.Run(`Backward knot returns /`, func(t *testing.T) {
		result, err := BackwardKnot.ToRune()

		checks.CheckHasNoError(t, result, err)
		if result != '/' {
			t.Errorf(`Expected /, got %v`, string(result))
		}
	})

	t.Run(`ForwardBackward knot returns >`, func(t *testing.T) {
		result, err := ForwardBackwardKnot.ToRune()

		checks.CheckHasNoError(t, result, err)
		if result != '>' {
			t.Errorf(`Expected >, got %v`, string(result))
		}
	})

	t.Run(`BackwardForward knot returns <`, func(t *testing.T) {
		result, err := BackwardForwardKnot.ToRune()

		checks.CheckHasNoError(t, result, err)
		if result != '<' {
			t.Errorf(`Expected <, got %v`, string(result))
		}
	})
}

func TestSwapsStrands(t *testing.T) {
	t.Run("Forward knot swaps strands", func(t *testing.T) {
		result := ForwardKnot.SwapsStrands()

		if !result {
			t.Errorf("expected true, got %t", result)
		}
	})

	t.Run("Backward knot swaps strands", func(t *testing.T) {
		result := BackwardKnot.SwapsStrands()

		if !result {
			t.Errorf("expected true, got %t", result)
		}
	})

	t.Run("ForwardBackward knot does not swap strands", func(t *testing.T) {
		result := ForwardBackwardKnot.SwapsStrands()

		if result {
			t.Errorf("expected false, got %t", result)
		}
	})

	t.Run("BackwardForward knot does not swap strands", func(t *testing.T) {
		result := BackwardForwardKnot.SwapsStrands()

		if result {
			t.Errorf("expected false, got %t", result)
		}
	})
}

func TestGetVisibleStrand(t *testing.T) {
	t.Run("Forward knot returns left strand", func(t *testing.T) {
		result := ForwardKnot.GetVisibleStrand()

		if result != LeftStrand {
			t.Errorf("Expected LeftStrand, got %v", result)
		}
	})

	t.Run("ForwardBackward knot returns left strand", func(t *testing.T) {
		result := ForwardBackwardKnot.GetVisibleStrand()

		if result != LeftStrand {
			t.Errorf("Expected RightStrand, got %v", result)
		}
	})

	t.Run("Backward knot returns right strand", func(t *testing.T) {
		result := BackwardKnot.GetVisibleStrand()

		if result != RightStrand {
			t.Errorf("Expected RightStrand, got %v", result)
		}
	})

	t.Run("BackwardForward knot returns right strand", func(t *testing.T) {
		result := BackwardForwardKnot.GetVisibleStrand()

		if result != RightStrand {
			t.Errorf("Expected RightStrand, got %v", result)
		}
	})
}

func TestParseKnots(t *testing.T) {
	t.Run("empty string returns empty slice", func(t *testing.T) {
		result, err := ParseKnots("")

		checks.CheckHasNoError(t, result, err)
		if len(result) != 0 {
			t.Errorf("Expected empty slice, got (%v, %v)", result, err)
		}
	})

	t.Run("Invalid character returns error", func(t *testing.T) {
		result, err := ParseKnots(">>🧶/")

		checks.CheckHasError(t, result, err, "unknown knot 🧶")
	})

	t.Run("Converts runes to the correct knots", func(t *testing.T) {
		result, err := ParseKnots(`>\<//`)

		checks.CheckHasNoError(t, result, err)

		expected := []Knot{
			ForwardBackwardKnot, ForwardKnot, BackwardForwardKnot, BackwardKnot, BackwardKnot,
		}

		if len(result) != len(expected) {
			t.Errorf("Wrong number of knots in result. actual=%v, expected=%v", result, expected)
		}

		for i, knot := range result {
			if knot != expected[i] {
				t.Errorf("Incorrect knot at position %d. actual=%v, expected=%v", i, result, expected)
			}
		}
	})
}
