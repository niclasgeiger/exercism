package sieve

const testVersion = 1

func Sieve(limit int) []int {
	marked := make(map[int]bool, limit+1)
	marked[0] = true
	marked[1] = true
	for i := 2; i <= limit; i++ {
		if !marked[i] {
			for j := 2 * i; j <= limit; j += i {
				marked[j] = true
			}
		}
	}
	out := []int{}
	for i := 2; i <= limit; i++ {
		if !marked[i] {
			out = append(out, i)
		}
	}
	return out
}
