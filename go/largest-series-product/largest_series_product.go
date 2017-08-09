package lsproduct

import (
	"errors"
	"strconv"
)

const testVersion = 5

func LargestSeriesProduct(digits string, span int) (int, error) {
	if len(digits) == 0 {
		if span == 0 {
			return 1, nil
		}
		return 0, errors.New("span is too big")
	}
	if span > len(digits) || span < 0 {
		return -1, errors.New("span is wrong value")
	}
	max, err := getProduct(digits[:span])
	if err != nil {
		return 0, err
	}
	for i := 1; i <= len(digits)-span; i++ {
		product, err := getProduct(digits[i : span+i])
		if err != nil {
			return 0, err
		}
		if product > max {
			max = product
		}
	}
	return max, nil
}

func getProduct(digits string) (int, error) {
	out := 1
	for _, r := range digits {
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			return 0, err
		}
		out *= digit
	}
	return out, nil
}
