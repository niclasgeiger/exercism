package cipher

import (
	"strings"
	"unicode"
)

const testVersion = 1

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

func GetLowerCaseNoSpace(i string) (out string) {
	for _, r := range i {
		if unicode.IsLetter(r) {
			out += strings.ToLower(string(r))
		}
	}
	return out
}
