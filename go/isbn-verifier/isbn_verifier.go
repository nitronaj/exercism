package isbn

func IsDigit(char rune) (int, bool) {
	if char < '0' || char > '9' {
		return 0, false
	}

	return int(char - '0'), true
}

func IsValidISBN(isbn string) bool {
	sum := 0
	pos := 10

	for _, char := range isbn {
		if digit, ok := IsDigit(char); ok {
			sum += digit * pos
			pos -= 1
		} else if pos == 1 && char == 'X' {
			sum += 10
			pos -= 1
		} else if char != '-' {
			return false
		}
	}

	return pos == 0 && sum%11 == 0
}
