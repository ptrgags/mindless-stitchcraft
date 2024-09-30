package knitting

import "errors"

// Check that a motif is nonempty and is composed of 'v' and '-' runes.
func ValidateMotif(motif string) error {
	if len(motif) == 0 {
		return errors.New("motif must not be empty")
	}

	for _, r := range motif {
		if r != 'v' && r != '-' {
			return errors.New("motif has invalid characters. It must be a string of knits ('v') and purls ('-')")
		}
	}

	return nil
}
