package atbash

import (
	"unicode"
	"strings"
)

const testVersion = 2

func Atbash(in string) string {
	out := ""
	count := 0
	for _, letter := range in {
		if unicode.IsLetter(letter) {
			letter = unicode.ToLower(letter)
			letter = rune('a' + ('z' - letter))
		}
		if unicode.IsLetter(letter) || unicode.IsNumber(letter) {
			out += string(letter)
			count++
			if count%5 == 0 {
				out += " "
			}
		}
	}
	return strings.TrimSpace(out)
}
