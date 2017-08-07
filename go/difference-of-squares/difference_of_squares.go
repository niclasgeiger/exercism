package diffsquares

import "math"

const testVersion = 1

func SquareOfSums(n int) int {
	return (int)(math.Pow(.5*(float64)(n)*(float64(n)+1), 2.0))
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
