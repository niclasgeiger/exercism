package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

func IsIsogram(s string) bool {
	alphabet := map[string]int{}
	for _, r := range s {
		if unicode.IsLetter(r) {
			letter := strings.ToLower(string(r))
			if alphabet[letter]++; alphabet[letter] > 1 {
				return false
			}
		}
	}
	return true
}
