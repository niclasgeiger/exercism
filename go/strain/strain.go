package strain

const testVersion = 1

type Ints []int
type Strings []string
type Lists [][]int

func (i Ints) Keep(predicate func(int) bool) Ints {
	var out Ints
	for _, ele := range i {
		if predicate(ele) {
			out = append(out, ele)
		}
	}
	return out
}

func (i Ints) Discard(predicate func(int) bool) Ints {
	var out Ints
	for _, ele := range i {
		if !predicate(ele) {
			out = append(out, ele)
		}
	}
	return out
}
func (i Strings) Keep(predicate func(string) bool) Strings {
	var out Strings
	for _, ele := range i {
		if predicate(ele) {
			out = append(out, ele)
		}
	}
	return out
}

func (s Strings) Discard(predicate func(string) bool) Strings {
	var out Strings
	for _, ele := range s {
		if !predicate(ele) {
			out = append(out, ele)
		}
	}
	return out
}

func (lists Lists) Keep(predicate func([]int) bool) Lists {
	var out Lists
	for _, ele := range lists {
		if predicate(ele) {
			out = append(out, ele)
		}
	}
	return out
}

func (lists Lists) Discard(predicate func([]int) bool) Lists {
	var out Lists
	for _, ele := range lists {
		if !predicate(ele) {
			out = append(out, ele)
		}
	}
	return out
}
