package bob

import (
	"strings"
	"unicode"
)

const testVersion = 3

func Hey(s string) string {
	s = strings.TrimSpace(s)
	if isEmpty(s) {
		return "Fine. Be that way!"
	}
	if isYelling(s) {
		return "Whoa, chill out!"
	}
	if isQuestion(s) {
		return "Sure."
	}
	return "Whatever."
}
func isEmpty(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsNumber(c) || c == '?' {
			return false
		}
	}
	return true
}

func isQuestion(s string) bool {
	return s[len(s)-1:] == "?"
}

func isYelling(s string) bool {
	onlyNumbers := true
	for _, c := range s {
		if unicode.IsLetter(c) || c == '!' {
			onlyNumbers = false
		}
	}
	if onlyNumbers {
		return false
	}
	for _, c := range s {
		if c == '!' {
			return true
		}
		if unicode.IsLetter(c) && unicode.IsLower(c) {
			return false
		}
	}
	return true
}
