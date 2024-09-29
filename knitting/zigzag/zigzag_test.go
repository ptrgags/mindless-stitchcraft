package zigzag

import (
	"testing"

	"github.com/ptrgags/mindless-stitchcraft/checks"
)

func TestGenerateZigzagPattern(t *testing.T) {
	t.Run("invalid motif returns error", func(t *testing.T) {
		validFabricWidth := 5
		cases := []struct {
			label         string
			motif         string
			expectedError string
		}{
			{"motif empty", "", "motif must not be empty"},
			{"motif has invalid characters", "--🧶vv", "motif has invalid characters"},
		}

		for _, tc := range cases {
			t.Run(tc.label, func(t *testing.T) {
				rows, err := GenerateZigzagPattern(tc.motif, validFabricWidth)

				checks.CheckHasError(t, rows, err, tc.expectedError)
			})
		}
	})

	t.Run("invalid fabricWidth returns error", func(t *testing.T) {
		validMotif := "v--v-"
		cases := []struct {
			label       string
			fabricWidth int
		}{
			{"zero", 0},
			{"negative", -1},
		}

		for _, tc := range cases {
			t.Run(tc.label, func(t *testing.T) {
				rows, err := GenerateZigzagPattern(validMotif, tc.fabricWidth)

				checks.CheckHasError(t, rows, err, "fabricWidth must be a positive integer")
			})
		}
	})

	t.Run("generated pattern is the correct shape", func(t *testing.T) {
		// The resulting pattern will always be:
		// width: fabricWidth
		// height: len(motif) / gcd(len(motif), fabricWidth)
		//    or twice this if there was an odd number of rows
		cases := []struct {
			label          string
			motif          string
			fabricWidth    int
			expectedHeight int
		}{
			{"len(motif) < fabricWidth, coprime widths", "v---", 5, 4},
			// height = 6 / gcd(6, 9) = 6 / 3 = 2
			{"len(motif) < fabricWidth, noncoprime widths", "---vv-", 9, 2},
			{"len(motif) > fabricWidth, coprime widths", "vvvv----", 3, 8},
			// height = 4 / gcd(4, 2) = 4 / 2 = 2
			{"len(motif) > fabricWidth, noncoprime widths", "----", 2, 2},
		}

		for _, tc := range cases {
			t.Run(tc.label, func(t *testing.T) {
				rows, err := GenerateZigzagPattern(tc.motif, tc.fabricWidth)

				checks.CheckHasNoError(t, rows, err)
				checks.CheckStringGridShape(t, rows, tc.fabricWidth, tc.expectedHeight)
			})
		}
	})

	t.Run("generated pattern is doubled to ensure even length", func(t *testing.T) {
		// The resulting pattern will always be:
		// width: fabricWidth
		// height: len(motif) / gcd(len(motif), fabricWidth)
		//    or twice this if there was an odd number of rows
		cases := []struct {
			label          string
			motif          string
			fabricWidth    int
			expectedHeight int
		}{
			// height 1 -> 2
			{"len(motif) == fabricWidth", "v--", 3, 2},
			// height 3 -> 6
			{"len(motif) < fabricWidth, coprime widths", "v--", 5, 6},
			// height = 5 / gcd(5, 10) = 5 / 5 = 1 -> 2
			{"len(motif) < fabricWidth, noncoprime widths", "v--v-", 10, 2},
			// height 5 -> 10
			{"len(motif) > fabricWidth, coprime widths", "-vvv-", 3, 10},
			// height = 9 / gcd(9, 3) = 9 / 3 = 3 -> 6
			{"len(motif) > fabricWidth, noncoprime widths", "--v---vvv", 3, 6},
		}

		for _, tc := range cases {
			t.Run(tc.label, func(t *testing.T) {
				rows, err := GenerateZigzagPattern(tc.motif, tc.fabricWidth)

				checks.CheckHasNoError(t, rows, err)
				checks.CheckStringGridShape(t, rows, tc.fabricWidth, tc.expectedHeight)
			})
		}
	})

	t.Run("generated pattern shows right side of fabric", func(t *testing.T) {
		cases := []struct {
			label        string
			motif        string
			fabricWidth  int
			expectedRows []string
		}{
			{"len(motif) < fabricWidth, coprime widths", "v---", 5, []string{
				"v-vvv",
				"--v--",
				"vvv-v",
				"v---v",
			}},
			// height = 6 / gcd(6, 9) = 6 / 3 = 2
			{"len(motif) < fabricWidth, noncoprime widths", "---vv-", 9, []string{
				"--vvvv--v",
				"----vv---",
			}},
			{"len(motif) > fabricWidth, coprime widths", "vvvv----", 3, []string{
				"vvv",
				"-vv",
				"v--",
				"---",
				"---",
				"v--",
				"-vv",
				"vvv",
			}},
			// height = 4 / gcd(4, 2) = 4 / 2 = 2
			{"len(motif) > fabricWidth, noncoprime widths", "----", 2, []string{
				"vv",
				"--",
			}},
		}
		for _, tc := range cases {
			t.Run(tc.label, func(t *testing.T) {
				rows, err := GenerateZigzagPattern(tc.motif, tc.fabricWidth)

				checks.CheckHasNoError(t, rows, err)
				checks.CheckStringGridsEqual(t, rows, tc.expectedRows)
			})
		}
	})

	t.Run("generated pattern shows right side of fabric (doubled cases)", func(t *testing.T) {
		cases := []struct {
			label        string
			motif        string
			fabricWidth  int
			expectedRows []string
		}{
			{"len(motif) == fabricWidth", "v--", 3, []string{
				"-vv",
				"--v",
			}},
			{"len(motif) < fabricWidth, coprime widths", "v--", 5, []string{
				"vv-vv",
				"v--v-",
				"-vv-v",
				"--v--",
				"v-vv-",
				"-v--v",
			}},
			{"len(motif) < fabricWidth, noncoprime widths", "v--v-", 10, []string{
				"-vv-v-vv-v",
				"-v--v-v--v",
			}},
			{"len(motif) > fabricWidth, coprime widths", "-vvv-", 3, []string{
				"--v",
				"v--",
				"---",
				"--v",
				"v--",
				"-vv",
				"vv-",
				"vvv",
				"-vv",
				"vv-",
			}},
			{"len(motif) > fabricWidth, noncoprime widths", "--v---vvv", 3, []string{
				"---",
				"---",
				"vv-",
				"vvv",
				"vvv",
				"v--",
			}},
		}

		for _, tc := range cases {
			t.Run(tc.label, func(t *testing.T) {
				rows, err := GenerateZigzagPattern(tc.motif, tc.fabricWidth)

				checks.CheckHasNoError(t, rows, err)
				checks.CheckStringGridsEqual(t, rows, tc.expectedRows)
			})
		}
	})
}
