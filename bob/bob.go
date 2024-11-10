package bob

import (
	"strings"
)

// Hey returns Bob's response to a remark
func Hey(remark string) string {
	// Trim leading and trailing spaces
	remark = strings.Trim(remark, "\n\r\t ")

	// check if remark is empty
	if remark == "" {
		return "Fine. Be that way!"
	}

	// get last character
	lastChar := string(remark[len(remark)-1])

	// check if remark is all uppercase
	remarkUpper := strings.ToUpper(remark)
	remarkLower := strings.ToLower(remark)

	// check if is all uppercase (yelling)
	isYelling := remark == remarkUpper && remark != remarkLower

	// check if is a question
	isQuestion := lastChar == "?"

	if isYelling {
		if isQuestion {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}

	// check if remark is a question
	if isQuestion {
		return "Sure."
	}

	return "Whatever."
}
