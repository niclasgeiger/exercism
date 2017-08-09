package palindrome

import (
	"errors"
	"math"
)

const testVersion = 1

func Products(min, max int) (minPalindrome Product, maxPalindrome Product, err error) {
	if min > max {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}
	minPalindrome = Product{Value: math.MaxInt32}
	maxPalindrome = Product{Value: math.MinInt32}
	var found bool
	for i := min; i <= max; i++ {
		for j := min; j <= max; j++ {
			product := Product{
				Value:          i * j,
				Factorizations: [][2]int{{i, j}},
			}
			if product.isPalindrome() {
				found = true
				if product.Value == minPalindrome.Value {
					minPalindrome = minPalindrome.Add(product)
				}
				if product.Value < minPalindrome.Value {
					minPalindrome = product
				}
				if product.Value == maxPalindrome.Value {
					maxPalindrome = maxPalindrome.Add(product)
				}
				if product.Value > maxPalindrome.Value {
					maxPalindrome = product
				}
			}
		}
	}
	if !found {
		return Product{}, Product{}, errors.New("no palindromes...")
	}
	return minPalindrome, maxPalindrome, nil
}
