package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}
	for key, value := range in {
		for _, letter := range value {
			out[strings.ToLower(letter)] = key
		}
	}
	return out
}
