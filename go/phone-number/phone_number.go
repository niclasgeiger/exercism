package phonenumber

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

const testVersion = 3

func Number(number string) (string, error) {
	if err := checkNumber(number); err != nil {
		return "", err
	}
	out := ""
	numbers := getNumbers(number)
	for _, num := range numbers {
		out += fmt.Sprintf("%d", num)
	}
	if len(out) == 11 {
		out = out[1:]
	}
	return out, nil
}

func AreaCode(number string) (string, error) {
	if err := checkNumber(number); err != nil {
		return "", err
	}
	bound := 0
	if countNumbers(number) == 11 {
		bound++
	}
	areaCode := ""
	numbers := getNumbers(number)
	for i := bound; i < bound+3; i++ {
		areaCode += fmt.Sprintf("%d", numbers[i])
	}
	return areaCode, nil
}

func Format(number string) (string, error) {
	numbers, err := Number(number)
	if err != nil {
		return "", err
	}
	if len(numbers) > 10 {
		numbers = numbers[1:]
	}
	out := fmt.Sprintf("(%s) %s-%s", numbers[:3], numbers[3:6], numbers[6:])
	return out, nil
}

func getNumbers(s string) []int {
	out := []int{}
	for _, r := range s {
		if unicode.IsNumber(r) {
			num, _ := strconv.Atoi(string(r))
			out = append(out, num)
		}
	}
	return out
}
func countNumbers(s string) int {
	count := 0
	for _, r := range s {
		if unicode.IsNumber(r) {
			count++
		}
	}
	return count
}
func checkNumber(number string) error {
	numbers := getNumbers(number)
	if countNumbers(number) == 10 {
		if numbers[0] < 2 || numbers[3] < 2 {
			return errors.New("first number in area code and exchange code  must be > 1")
		}
		return nil
	}
	if countNumbers(number) == 11 {
		if numbers[1] < 2 || numbers[4] < 2 {
			return errors.New("first number in area code and exchange code  must be > 1")
		}
		if number[:2] == "+1" {
			return nil
		}
		if number[0] == '1' {
			return nil
		}
	}

	return errors.New("Bad Format")
}
