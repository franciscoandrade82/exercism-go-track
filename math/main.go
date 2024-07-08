package main

import (
	"fmt"
	"math"
)

func main() {
	
	i1, i2, i3 := 1, 2, 3
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 2.1, 3.2, 4.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", math.Round(floatSum*100)/100)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi

	fmt.Printf("Circumference: %.2f\n", circumference)
}
