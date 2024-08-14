package patterns

import (
	"errors"
	"regexp"
	"strings"
)

func validateStitches(stitches string) error {
	regex := regexp.MustCompile("^(v|-)*$")
	if stitches == "" || !regex.MatchString(stitches) {
		return errors.New("invalid stitches string. It must be a string of at least 1 character composed of knits ('v') and purls ('-')")
	}

	return nil
}

// Take a string of stiches (either 'v' for knit or '-' for purl)
// and swap the knits and purls. This is one part of what happens when
// you flip the fabric over when knitting. This returns a new string
func SwapKnitsAndPurls(stitches string) (string, error) {
	if err := validateStitches(stitches); err != nil {
		return "", err
	}

	return strings.Map(func(r rune) rune {
		if r == 'v' {
			return '-'
		} else if r == '-' {
			return 'v'
		}

		return r
	}, stitches), nil
}

func generatePatternShort(stitches string, width int) ([]string, error) {
	return []string{}, nil
}

func generatePatternLong(stitches string, width int) ([]string, error) {
	return []string{}, nil
}

// Take a pattern of knits and purls e.g. repeat it
// over and over until the pattern repeats for the given number of stitches
// wide. E.g. for input "v--", 5 we have:
//
// v--v-
// -v--v
// --v--
//
// This method assumes all stitches are on the right side of the fabric,
// and the first stitch is the top left, other functions will convert this
// to a usable knitting chart
func GeneratePattern(stitches string, width int) ([]string, error) {
	if err := validateStitches(stitches); err != nil {
		return []string{}, err
	}

	n := len(stitches)

	// Handle simple cases
	if n == width {
		return []string{stitches}, nil
	}

	// The logic is a little different depending on whether the stitches
	// fit within the width or the other way around.
	if n < width {
		return generatePatternShort(stitches, width)
	}

	return generatePatternLong(stitches, width)
}
