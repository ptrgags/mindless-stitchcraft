package patterns

import (
	"errors"
	"regexp"
	"strings"
)

// Take a string of stiches (either 'v' for knit or '-' for purl)
// and swap the knits and purls. This is one part of what happens when
// you flip the fabric over when knitting. This returns a new string
func SwapKnitsAndPurls(stitches string) (string, error) {
	regex := regexp.MustCompile("^(v|-)*$")
	if !regex.MatchString(stitches) {
		return "", errors.New("invalid stitches string. Stitches are composed of knits ('v') and purls ('-')")
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
