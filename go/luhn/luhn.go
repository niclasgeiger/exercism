package luhn

import (
	"sort"
	"strconv"
	"unicode"
)

const testVersion = 2

func Valid(num string) bool {
	var sorted sort.StringSlice
	for i := len(num) - 1; i >= 0; i-- {
		r := rune(num[i])
		if unicode.IsNumber(r) {
			sorted = append(sorted, string(r))
		} else if !unicode.IsSpace(r) {
			return false
		}
	}

	i := 0
	sum := 0
	for _, s := range sorted {
		digit, _ := strconv.Atoi(s)
		if i%2 == 1 {
			digit = (digit * 2)
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		i++
	}
	if sum == 0 && len(sorted) < 2 {
		return false
	}
	return sum%10 == 0
}
