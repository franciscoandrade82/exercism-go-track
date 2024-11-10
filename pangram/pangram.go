package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	seen := 0
	for _, r := range strings.ToLower(input) {
		if 'a' <= r && r <= 'z' {
			// Shift the rune to get a 0-25 index for each letter.
			idx := int(r - 'a')
			// Use bitwise OR to mark the letter as seen.
			seen |= 1 << idx
		}
	}

	// check if all 26 letters were seen
	return seen == (1<<26)-1
}
