package diamond

import (
	"errors"
)

const testVersion = 1

var alphabet []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Gen(letter byte) (string, error) {
	pos := letterPosition(letter)
	if pos < 0 {
		return "", errors.New("no valid letter")
	}
	out := ""
	for i := 0; i < pos; i++ {
		out += drawLine(pos, i)
	}
	for i := pos - 2; i >= 0; i-- {
		out += drawLine(pos, i)
	}
	return out, nil
}

func letterPosition(s byte) int {
	for index, letter := range alphabet {
		if byte(s) == letter {
			return index + 1
		}
	}
	return -1
}

func drawLine(pos, i int) string {
	line := ""
	for j := 0; j < pos-i-1; j++ {
		line += " "
	}
	line += string(alphabet[i])
	for j := 0; j < 2*i-1; j++ {
		line += " "
	}
	if i > 0 {
		line += string(alphabet[i])
	}
	for j := 0; j < pos-i-1; j++ {
		line += " "
	}
	return line + "\n"
}
