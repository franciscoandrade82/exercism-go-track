package main

import (
	"fmt"
	"sort"
)

func main() {
	colors := [3]string{"Red", "Green", "Blue"} // an array of strings
	newColors := []string{"Pink", "Purple", "Magenta"} // a slice of strings

	fmt.Println(colors)
	fmt.Println(newColors)

	newColors = append(newColors, "Violet")
	fmt.Println(newColors)

	fmt.Printf("Colors: %T\n", colors)
	fmt.Printf("New Colors: %T\n", newColors)

	numbers := make([]int, 5)
	numbers[0] = 8
	numbers[1] = 2
	numbers[2] = 6
	numbers[3] = 4
	numbers[4] = 5

	numbers = append(numbers, 6)
	
	fmt.Println(numbers)

	sort.Ints(numbers)

	fmt.Println(numbers)
}
