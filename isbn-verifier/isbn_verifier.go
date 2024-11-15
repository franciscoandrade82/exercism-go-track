package isbn

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	// clean dashes from ISBN
	isbn = strings.ReplaceAll(isbn, "-", "")

	// check lenght
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i, r := range isbn[:9] {
		if !unicode.IsDigit(r) {
			return false
		}

		sum += int(r-'0') * (10 - i)
	}

	// validate the check digit
	checkDigit := isbn[9]
	if checkDigit == 'X' || checkDigit == 'x' {
		sum += 10
	} else if unicode.IsDigit(rune(checkDigit)) {
		sum += int(checkDigit - '0')
	} else {
		return false
	}

	return sum%11 == 0
}
