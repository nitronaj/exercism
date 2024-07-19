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
	var sum uint64
	boardSquares := 64
	for i := 1; i <= boardSquares; i++ {
		grains, err := Square(i)
		if err != nil {
			return 0
		}
		sum += grains
	}
	return sum
}
