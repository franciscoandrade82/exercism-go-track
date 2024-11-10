// Write your answer here, and then test your code.
// Your job is to implement the getCartFromJson() method.

package main

import (
	"encoding/json"
	"fmt"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

type cartItem struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []cartItem {
	var cart []cartItem
	err := json.Unmarshal([]byte(jsonString), &cart)
	if err != nil {
		panic(err)
	}

	return cart
}

func main() {
	jsonString :=
		`[{"name":"apple","price":2.99,"quantity":3},
	  {"name":"orange","price":1.50,"quantity":8},
	  {"name":"banana","price":0.49,"quantity":12}]`

	result := getCartFromJson(jsonString)

	for _, item := range result {
		fmt.Println(item.Name)
	}
}
