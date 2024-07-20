package isbn

import (
	"unicode"
)

func IsValidISBN(isbn string) bool {
	sum := 0
	multiply := 10
	countDigits := 0
	for _, char := range isbn {
		if unicode.IsDigit(char) {
			sum += int(char-'0') * multiply
			multiply -= 1
			countDigits += 1
		}

		if unicode.IsLetter(char) {
			if char == 'X' {
				sum += 10
				countDigits += 1
			} else {
				return false
			}
		}
	}

	return countDigits == 10 && sum%11 == 0
}
