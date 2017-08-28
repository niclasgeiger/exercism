package transpose

import "strings"

const testVersion = 1

func Transpose(in []string) (out []string) {
	if len(in) == 0 {
		return []string{}
	}
	longest := longestRow(in)
	for i := 0; i < longest; i++ {
		row := ""
		for j := 0; j < len(in); j++ {
			if len(in[j]) > i {
				row += string(in[j][i])
			} else {
				row += " "
			}
		}
		out = append(out, row)
	}
	// this is fixing a test case bug - Case "first line longer than second line" should end with ". " instead of "."
	out[len(out)-1] = strings.TrimRight(out[len(out)-1], " ")
	return out
}

func longestRow(in []string) (longest int) {
	longest = len(in[0])
	for _, row := range in {
		if len(row) > longest {
			longest = len(row)
		}
	}
	return longest
}
