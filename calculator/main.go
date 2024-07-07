// Write your answer here, and then test your code.

package calculator

import "strconv"

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
    // Your code goes here.
	var result float64

	value1 := convertInputToValue(input1)
	value2 := convertInputToValue(input2)

	switch operation {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subtractValues(value1, value2)
	case "*":
		result = multiplyValues(value1, value2)
	case "/":
		result = divideValues(value1, value2)
	default:
		panic("Invalid operation")
	}

	return result
}

func convertInputToValue(input string) float64 {
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err)
	}
	
	return value
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}

func main() {
	// This function is not used in the test.
	calculate("1", "2", "+")
	calculate("1", "2", "-")
	calculate("1", "2", "*")
	calculate("1", "2", "/")
}