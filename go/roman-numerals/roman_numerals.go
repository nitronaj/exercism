package romannumerals

import (
	"fmt"
	"strings"
)

//	M	D	C	L	X	V	I
//
// 1000	500	100	50	10	5	1

type RomanNumeral struct {
	roman   string
	decimal int
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", fmt.Errorf("wrong input %d", input)
	}

	var result string
	romanNumerals := []RomanNumeral{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"I", 1},
	}

	for i, romanNumeral := range romanNumerals {
		roman, decimal := romanNumeral.roman, romanNumeral.decimal
		d := input / decimal

		if decimal == 1000 {
			result += strings.Repeat(roman, d)
		} else {
			if d == 9 {
				hRoman := romanNumerals[i-2].roman
				result += roman + hRoman
			} else if d >= 5 {
				c := d - 5
				hRoman := romanNumerals[i-1].roman
				result += hRoman + strings.Repeat(roman, c)
			} else if d == 4 {
				hRoman := romanNumerals[i-1].roman
				result += roman + hRoman
			} else {
				result += strings.Repeat(roman, d)
			}
		}
		input = input % decimal
	}

	fmt.Println(result)

	return result, nil
}
