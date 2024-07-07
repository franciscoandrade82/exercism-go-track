package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current time:", now)

	birthday := time.Date(1982, 2, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Birthday:", birthday)

	fmt.Println(birthday.Format(time.ANSIC))
}
