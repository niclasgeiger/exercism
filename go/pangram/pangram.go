package pangram

import (
	"unicode"
)

const testVersion = 1

func IsPangram(s string) bool {
	if len(s) < 26 {
		return false
	}
	alphabet := map[string]int{}
	for _, r := range s {
		if unicode.Is(unicode.Latin, r) {
			letter := string(r)
			alphabet[letter]++
		}
	}
	return len(alphabet) > 25
}
