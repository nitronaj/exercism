// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob provides utility functions and types for handling various operations related to "bob" resonse.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case isQuestion(remark) && isYelling(remark):
		return "Calm down, I know what I'm doing!" // This is what he says if you yell a question at him.
	case isQuestion(remark):
		return "Sure." //   The convention used for questions is that it ends with a question mark.
	case isYelling(remark): // The convention used for yelling is ALL CAPITAL LETTERS.
		return "Whoa, chill out!"
	case isSilence(remark):
		return "Fine. Be that way!" // This is how he responds to silence. The convention used for silence is nothing, or various combinations of whitespace characters.
	default:
		return "Whatever."
	}
}

func isQuestion(str string) bool {
	return str != "" && string(str[len(str)-1]) == "?"
}

func isYelling(str string) bool {
	hasAlphabet := false
	for _, char := range str {
		if unicode.IsLetter(char) {
			if !unicode.IsUpper(char) {
				return false
			}

			hasAlphabet = true
		}
	}

	return hasAlphabet
}

func isSilence(str string) bool {
	return len(str) == 0

}
