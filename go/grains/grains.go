package grains

import (
	"errors"
)

const testVersion = 1

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("parameter out of interval")
	}
	return pow(n), nil
}

//total is 2^65 - 1
func Total() uint64 {
	return pow(65) - 1
}

//2^(x-1) per square
func pow(n int) uint64 {
	return 1 << (uint64)(n-1)
}
