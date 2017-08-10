package wordcount

import (
	"strings"
	"unicode"
)

const testVersion = 3

// Use this return type.
type Frequency map[string]int

// Just implement the function.
func WordCount(phrase string) Frequency {
	words := split(phrase)
	out := map[string]int{}
	for _, word := range words {
		out[word]++
	}
	return out
}

func split(s string) []string {
	s = strings.TrimSpace(s)
	word := ""
	out := []string{}
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '\'' {
			word += strings.ToLower(string(r))
		} else if unicode.IsSpace(r) || unicode.IsPunct(r) {
			if len(word) > 0 {
				word = strings.Trim(word, "'")
				out = append(out, word)
				word = ""
			}
		}
	}
	if len(word) > 0 {
		word = strings.Trim(word, "'")
		out = append(out, word)
	}
	return out
}
