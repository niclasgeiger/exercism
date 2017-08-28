package prime

const (
	testVersion = 1
	CEIL        = 105000 // this depends on the number of primes we want to generate
)

var sieve = sieveOfSundaram(CEIL)

func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	if len(sieve)+1 < n {
		return 0, false
	}
	return sieve[n-1], true
}

func sieveOfSundaram(limit int) (out []int) {
	n := limit / 2
	primes := make([]bool, limit)
	for i := 1; i < n; i++ {
		for j := 1; j <= (n-i)/(2*i+1); j++ {
			primes[i+j+2*i*j] = true
		}
	}
	out = []int{2, 3}
	for i := 2; i < len(primes)/2; i++ {
		if !primes[i] {
			out = append(out, 2*i+1)
		}
	}
	return out
}
