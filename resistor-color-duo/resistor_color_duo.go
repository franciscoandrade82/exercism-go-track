package resistorcolorduo

// create a map of colors to values
var colorValues = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	if len(colors) < 2 {
		return -1
	}

	// get the value of the first two colors
	return colorValues[colors[0]]*10 + colorValues[colors[1]]
}
