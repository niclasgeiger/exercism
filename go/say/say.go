package say

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const testVersion = 1

func Say(num uint64) string {
	if num == 0 {
		return "zero"
	}
	number := fmt.Sprintf("%d", num)
	numbers, err := explode(number)
	if err != nil {
		return ""
	}
	out := []string{}
	for i, n := range numbers {
		said := say(n)
		if said != "" {
			out = append(out, said)
			step := steps[len(numbers)-i-1]
			out = append(out, step)
		}
	}
	return strings.TrimSpace(strings.Join(out, " "))
}

func explode(number string) ([]int, error) {
	numbers := []int{}
	if off := len(number) % 3; off > 0 {
		n, err := strconv.Atoi(number[:off])
		if err != nil {
			return []int{}, errors.New("error parsing")
		}
		numbers = append(numbers, n)
		number = number[off:]
	}
	for i := 0; i <= len(number)-3; i += 3 {
		n, err := strconv.Atoi(number[i : i+3])
		if err != nil {
			return []int{}, errors.New("error parsing")
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}

func say(num int) string {
	words := []string{}
	if num >= 100 {
		words = append(words, spelling[num/100])
		words = append(words, "hundred")
		num = num % 100
	}
	if num >= 20 {
		s := tens[num-(num%10)]
		num -= num - (num % 10)
		if num > 0 {
			s += "-" + spelling[num]
			num = 0
		}
		words = append(words, s)
	}
	if num > 10 {
		words = append(words, teens[num])
		num = 0
	}
	if num > 0 {
		words = append(words, spelling[num])
	}
	return strings.Join(words, " ")
}
