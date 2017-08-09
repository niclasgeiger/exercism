package palindrome

import "strconv"

type Product struct {
	Value          int
	Factorizations [][2]int
}

func (p Product) Add(a Product) Product {
	if !p.containsFactors(a.Factorizations[0]) {
		p.Factorizations = append(p.Factorizations, a.Factorizations...)
	}
	return p
}

func (p Product) containsFactors(comp [2]int) bool {
	for _, fac := range p.Factorizations {
		if (fac[0] == comp[0] && fac[1] == comp[1]) || (fac[0] == comp[1] && fac[1] == comp[0]) {
			return true
		}
	}
	return false
}

func (p Product) isPalindrome() bool {
	// no negative product is a palindrome!
	if p.Value < 0 {
		return false
	}
	num := strconv.Itoa(p.Value)
	for i := 0; i < len(num)/2; i++ {
		if num[i] != num[len(num)-1-i] {
			return false
		}
	}
	return true
}
