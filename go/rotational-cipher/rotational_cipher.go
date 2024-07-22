package rotationalcipher

import (
	"unicode"
)

func ShiftChar(char rune, shiftKey int) rune {
	newChar := int(char) + shiftKey

	if unicode.IsLower(char) && int(newChar) > 122 {
		newChar = newChar - int('z') - 1 + int('a')
	} else if unicode.IsUpper(char) && int(newChar) > 90 {
		newChar = newChar - int('Z') - 1 + int('A')
	}
	return rune(newChar)
}

func RotationalCipher(plain string, shiftKey int) string {
	var result string
	for _, char := range plain {
		if unicode.IsLetter(char) {
			result += string(ShiftChar(char, shiftKey))
		} else {
			result += string(char)
		}
	}

	return result
}
