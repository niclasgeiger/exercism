package etl

import "strings"

const testVersion = 1

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}
	for score, letters := range in {
		for _, letter := range letters {
			out[strings.ToLower(letter)] = score
		}
	}
	return out
}
