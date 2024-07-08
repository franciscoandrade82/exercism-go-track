package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello,", name)

	fmt.Print("Enter your age: ")
	age, _ := reader.ReadString('\n')

	ageInt, err := strconv.ParseInt(strings.TrimSpace(age), 10, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("You are", ageInt, "years old.")
}
