package resistorcolorduo

import "math"

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {

	mapColors := map[string]int{
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
	resistance := 0

	for i := 0; i < len(colors) && i < 2; i++ {
		resistance += mapColors[colors[i]] * int(math.Pow10(1-i))
	}

	return resistance
}
