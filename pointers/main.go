package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p = &anInt // p is a pointer to anInt, p stores the memory address of anInt
	fmt.Println("anInt:", anInt)
	fmt.Println("p:", p)
	fmt.Println("*p:", *p) // *p is the value stored at the memory address p
}
