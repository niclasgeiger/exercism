package summultiples

const testVersion = 2

func SumMultiples(limit int, divisors ...int) int {
	out := 0
	for i := 1; i < limit; i++ {
		if isDivisor(i, divisors) {
			out += i
		}
	}
	return out
}

func isDivisor(num int, divisors []int) bool {
	for _, div := range divisors {
		if num%div == 0 {
			return true
		}
	}
	return false
}
