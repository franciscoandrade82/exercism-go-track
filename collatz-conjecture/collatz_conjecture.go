package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("the number must be a positive number")
	}

	return collatzSteps(n), nil
}

func collatzSteps(n int) int {
	if n == 1 {
		return 0
	}

	if n%2 == 0 {
		return 1 + collatzSteps(n/2)
	}

	return 1 + collatzSteps(3*n+1)
}
