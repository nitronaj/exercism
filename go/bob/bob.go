// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob provides utility functions and types for handling various operations related to "bob" resonse.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

type Remark string

func NewRemark(remark string) Remark {
	return Remark(strings.TrimSpace(remark))
}

func (r *Remark) String() string {
	return string(*r)
}
func (r *Remark) isQuestion() bool {
	return strings.HasSuffix(r.String(), "?")
}

func (r *Remark) isYelling() bool {
	hasAlphabet := strings.IndexFunc(r.String(), unicode.IsLetter) >= 0
	isUpperCase := strings.ToUpper(r.String()) == r.String()

	return hasAlphabet && isUpperCase
}

func (r *Remark) isYellingQuestion() bool {
	return r.isYelling() && r.isQuestion()
}

func (r *Remark) isSilence() bool {
	return *r == ""
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	r := NewRemark(remark)

	switch {
	case r.isYellingQuestion():
		return "Calm down, I know what I'm doing!" // This is what he says if you yell a question at him.
	case r.isQuestion():
		return "Sure." //   The convention used for questions is that it ends with a question mark.
	case r.isYelling(): // The convention used for yelling is ALL CAPITAL LETTERS.
		return "Whoa, chill out!"
	case r.isSilence():
		return "Fine. Be that way!" // This is how he responds to silence. The convention used for silence is nothing, or various combinations of whitespace characters.
	default:
		return "Whatever."
	}
}
