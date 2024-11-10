package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {

	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	var sum int
	for i := len(id) - 1; i >= 0; i-- {
		char := string(id[i])
		n, err := strconv.Atoi(char)
		if err != nil {
			return false
		}

		if (len(id)-1-i)%2 != 0 {
			n = n * 2
			if n > 9 {
				n -= 9
			}
		}

		sum += n
	}

	return sum%10 == 0
}
