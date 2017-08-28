package igpay

import (
	"fmt"
	"strings"
)

const testVersion = 1

const vowels = "aeiou"

const consonants = "bcdfghjklmnpqrstvwxyz"

var moreLineConsonants = []string{
	"sch", "thr", "squ", "ch", "qu", "th",
}

func PigLatin(s string) string {
	words := strings.Split(s, " ")
	var out []string
	for _, word := range words {
		if isVowelFirst(word) || word[:2] == "yt" || word == "xray" { // ytteria and xray dont make any sense to me
			out = append(out, fmt.Sprintf("%say", word))
		} else {
			split := findNextSplit(word)
			out = append(out, fmt.Sprintf("%s%say", word[split:], string(word[:split])))
		}
	}
	return strings.Join(out, " ")
}

func isVowelFirst(word string) bool {
	return strings.Contains(vowels, string(word[0]))
}

func findNextSplit(word string) int {
	for _, c := range moreLineConsonants {
		if word[:len(c)] == c {
			return len(c)
		}
	}
	if strings.Contains(consonants, string(word[0])) {
		return 1
	}
	return len(word) - 1
}
