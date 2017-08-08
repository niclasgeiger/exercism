package cryptosquare

import (
	"bytes"
	"math"
	"strings"
	"unicode"
)

const testVersion = 2

func Encode(s string) string {
	transformed := trim(s)
	c := getBounds(transformed)
	square := getSquare(c, transformed)
	return encrypt(c, square)
}

func getBounds(s string) int {
	r := int(math.Sqrt(float64(len(s))))
	c := r
	if r*r < len(s) {
		c++
	}
	return c
}

func encrypt(c int, square []string) string {
	encrypted := []string{}
	for i := 0; i < c; i++ {
		var encryptedLineBuffer bytes.Buffer
		for _, line := range square {
			if i < len(line) {
				encryptedLineBuffer.Write([]byte{line[i]})
			}
		}
		encrypted = append(encrypted, encryptedLineBuffer.String())
	}
	return strings.Join(encrypted, " ")
}

func getSquare(c int, s string) []string {
	square := []string{}
	for i := 0; i < len(s); i += c {
		end := i + c
		if len(s) < i+c {
			end = len(s)
		}
		square = append(square, s[i:end])
	}
	return square
}

func trim(s string) string {
	var buffer bytes.Buffer
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			buffer.Write([]byte(strings.ToLower(string(r))))
		}
	}
	return buffer.String()
}
