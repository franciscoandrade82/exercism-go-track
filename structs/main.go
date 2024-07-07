package main

import (
	"fmt"
)

type Dog struct {
	Name string
	Age int
	width int
}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

func main() {
	dog := Dog{
		Name: "Fido",
		Age: 3,
		width: 10,
	}
	
	fmt.Println(dog)
	fmt.Println(dog.Name)

	dog.Name = "Rover"
	fmt.Println(dog.Name)
	
	fmt.Printf("%+v\n", dog)

	dog.Speak()
}

