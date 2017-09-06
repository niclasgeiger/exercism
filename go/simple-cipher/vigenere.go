package cipher

import (
	"unicode"
)

type VigenereKey []rune

type Vigenere struct {
	key VigenereKey
}

func NewVigenere(key string) *Vigenere {
	if len(key) < 3 {
		return nil
	}
	var k = VigenereKey{}
	for _, r := range key {
		if !unicode.IsLetter(r) || unicode.IsUpper(r) {
			return nil
		}
		r = unicode.ToLower(r)
		k = append(k, rune(r-97))
	}
	return &Vigenere{
		key: k,
	}
}

func (vigenere Vigenere) Encode(encoded string) string {
	remove := func(r rune, i int) rune {
		return rune(r + vigenere.key[i%len(vigenere.key)])
	}
	return vignereHelp(encoded, remove)
}

func (vigenere Vigenere) Decode(clear string) string {
	add := func(r rune, i int) rune {
		return rune(r - vigenere.key[i%len(vigenere.key)])
	}
	return vignereHelp(clear, add)
}

func vignereHelp(v string, f func(rune, int) rune) (out string) {
	count := 0
	for _, r := range v {
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
			next := f(r, count)
			if next < 97 {
				next += 26
			}
			if next > 122 {
				next -= 26
			}
			out += string(next)
			count++
		}
	}
	return out
}
