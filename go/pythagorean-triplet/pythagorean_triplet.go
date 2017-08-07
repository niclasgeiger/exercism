package pythagorean

const testVersion = 1

/*
	Using Dicksons Method
	x = r + s
	y = r + t
	z = r + s + t
	r^2 = 2st
	n = x + y + z = 3r + 2s + 2t
*/

type Triplet struct {
	A int
	B int
	C int
}

func Sum(size int) []Triplet {
	triplets := []Triplet{}
	for i := 1; i < size/3; i++ {
		found := getTrianglePermutations(i, size)
		if found != nil {
			triplets = append(triplets, *found)
		}
	}
	return triplets
}
func getTrianglePermutations(r, size int) *Triplet {
	if (r*r)%2 > 0 {
		return nil
	}
	st := r * r / 2
	for s := 1; s < r; s++ {
		if st%s == 0 {
			t := st / s
			triplet := &Triplet{
				A: r + s,
				B: r + t,
				C: r + s + t,
			}
			if triplet.Size() == size {
				return triplet
			}
		}
	}
	return nil
}

func getTriangles(r, min, max int) []Triplet {
	out := []Triplet{}
	if (r*r)%2 > 0 {
		return out
	}
	st := r * r / 2
	for s := 1; s < r; s++ {
		if st%s == 0 {
			t := st / s
			triplet := Triplet{
				A: r + s,
				B: r + t,
				C: r + s + t,
			}
			if triplet.A >= min && triplet.B >= min && triplet.C <= max {
				out = append(out, triplet)
			}
		}
	}
	return out
}

func Range(start, end int) []Triplet {
	out := []Triplet{}
	for i := 1; i <= end; i++ {
		out = append(out, getTriangles(i, start, end)...)
	}
	return out
}

func (t Triplet) Size() int {
	return t.A + t.B + t.C
}
