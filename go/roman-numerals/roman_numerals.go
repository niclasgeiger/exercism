package romannumerals

import "errors"

const testVersion = 3

type RomanMapping struct {
	num     int
	literal string
}

var romanMappings = []RomanMapping{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(arabic int) (string, error) {
	if arabic < 1 || arabic >= 4000 {
		return "", errors.New("only numbers in between 1 and 3999 permitted")
	}
	out := ""
	for _, mapping := range romanMappings {
		for arabic >= mapping.num {
			out += mapping.literal
			arabic -= mapping.num
		}
	}
	return out, nil
}
