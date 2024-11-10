package romannumerals

import "fmt"

type arabicToRomanMapper struct {
	arabic int
	roman  string
}

var arabicToRoman = []arabicToRomanMapper{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", fmt.Errorf("number is out of range")
	}

	var number string
	for _, mapper := range arabicToRoman {
		count := input / mapper.arabic
		if count > 0 {
			number += repeat(mapper.roman, count)
			input -= mapper.arabic * count
		}

		if input == 0 {
			break
		}
	}

	return number, nil
}

// Helper function to repeat a string `count` times
func repeat(roman string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += roman
	}
	return result
}
