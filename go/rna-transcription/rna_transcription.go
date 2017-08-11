package strand

import "strings"

const testVersion = 3

func mapping(r rune) rune {
	switch r {

	case 'G':
		return 'C'

	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	}
	return '-'
}
func ToRNA(s string) string {
	return strings.Map(mapping, s)
}
