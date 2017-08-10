package allergies

const testVersion = 1

var allAllergies = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

func Allergies(n uint) []string {
	out := []string{}
	for i := uint(2048); i > 0; i = i >> 1 {
		if n < i {
			continue
		}
		n = n - i
		if i <= 128 {
			out = append(out, allAllergies[i])
		}
	}
	return out
}

func AllergicTo(n uint, allergen string) bool {
	person := Allergies(n)
	for _, a := range person {
		if a == allergen {
			return true
		}
	}
	return false
}
