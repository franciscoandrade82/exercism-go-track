package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	count := make(map[rune]int)

	for _, l := range strings.ToLower(word) {
		if unicode.IsLetter(l) {
			count[l]++
			if count[l] > 1 {
				return false
			}
		}
	}

	return true
}
