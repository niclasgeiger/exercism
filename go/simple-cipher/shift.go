package cipher

import "unicode"

type Shift struct {
	s int
}

func NewShift(i int) *Shift {
	if i < -25 || i > 25 || i == 0 {
		return nil
	}
	return &Shift{
		s: i,
	}
}

func (shift Shift) Encode(clear string) (out string) {
	return shiftHelp(clear, func(r rune) rune { return r + int32(shift.s) })
}

func (shift Shift) Decode(encoded string) (out string) {
	return shiftHelp(encoded, func(r rune) rune { return r - int32(shift.s) })
}

func shiftHelp(shift string, f func(rune) rune) (out string) {
	for _, r := range shift {
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
			next := f(r)
			if next < 97 {
				next += 26
			}
			if next > 122 {
				next -= 26
			}
			out += string(next)
		}
	}
	return out
}
