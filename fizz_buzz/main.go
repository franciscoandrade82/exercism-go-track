package main

import (
	"fmt"
)

/*
if the number is divisible by three, let's say nine, you need to print fizz instead of the number. If the number is visible by five, let's say 10, you need to print buzz instead of the number. If the number is visible by both three and five, let's say 15, you need to print fizz buzz. And otherwise you just print the number. This is the statement definition and here is the way to check if a number is divisible other number by using the modulo. So one modulo five will bring one, seven modulo five will print out two, this is the remainder, and 10 modulo five is going to be zero.
*/
func main() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}