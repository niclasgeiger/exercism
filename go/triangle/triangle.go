package triangle

import "math"

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	k := Kind{a, b, c}
	if !k.isTriangle() {
		return NaT
	}
	if k.isEquilateral() {
		return Equ
	}
	if k.isIsosceles() {
		return Iso
	}
	return Sca
}

func (k Kind) isEquilateral() bool {
	return k.A == k.B && k.B == k.C && k.A == k.C
}
func (k Kind) isIsosceles() bool {
	return k.A == k.B || k.B == k.C || k.A == k.C
}
func (k Kind) isTriangle() bool {
	if k.A == math.Inf(1) || k.A == math.Inf(-1) || k.B == math.Inf(1) || k.B == math.Inf(-1) || k.C == math.Inf(1) || k.C == math.Inf(-1) {
		return false
	}
	a := k.B + k.C
	b := k.A + k.C
	c := k.A + k.B
	return a >= k.A && b >= k.B && c >= k.C && a > 0 && b > 0 && c > 0
}

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind struct {
	A float64
	B float64
	C float64
}

var (
	NaT = Kind{0, 0, 0} // not a triangle
	Equ = Kind{1, 1, 1} // equilateral
	Iso = Kind{1, 2, 2} // isosceles
	Sca = Kind{3, 4, 5} // scalene
)
