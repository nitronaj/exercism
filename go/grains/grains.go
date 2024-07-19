package grains

import (
	"fmt"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, fmt.Errorf("square number must be between 1 and 64, got %d", number)
	}

	return 1 << (number - 1), nil
}

func Total() uint64 {
	return 1<<64 - 1
}
