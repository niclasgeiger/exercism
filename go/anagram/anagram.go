package anagram

import (
	"strings"
	"unicode"
)

const testVersion = 2

func Detect(word string, candidates []string) []string {
	counts := map[string]int{}
	out := []string{}
	rangeOver(word, counts, countUp)
	for _, candidate := range candidates {
		countsCopy := copy(counts)
		if len(candidate) != len(word) || strings.ToLower(word) == strings.ToLower(candidate) {
			continue
		}
		if rangeOver(candidate, countsCopy, countDown) {
			out = append(out, candidate)
		}
	}
	return out
}

func countDown(letter string, counts map[string]int) bool {
	if counts[letter]--; counts[letter] < 0 {
		return false
	}
	return true
}

func countUp(letter string, counts map[string]int) bool {
	counts[letter]++
	return true
}

func rangeOver(s string, count map[string]int, f func(string, map[string]int) bool) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			letter := strings.ToLower(string(r))
			if !f(letter, count) {
				return false
			}
		}
	}
	return true
}

func copy(m map[string]int) map[string]int {
	copy := map[string]int{}
	for k, v := range m {
		copy[k] = v
	}
	return copy
}
