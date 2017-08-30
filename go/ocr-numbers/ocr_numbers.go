package ocr

import (
	"fmt"
	"strings"
)

const testVersion = 1

func recognizeDigit(digit string) string {
	rawInput := strings.Replace(digit, "\n", "", -1)
	for i := 0; i < 10; i++ {
		rawData := strings.Replace(digitData[i], "\n", "", -1)
		if rawData == rawInput {
			return fmt.Sprintf("%d", i)
		}
	}
	return "?"
}

func Recognize(digits string) (out []string) {
	lines := getLines(digits)
	for _, line := range lines {
		ocrLine := ""
		for i := 0; i < len(line)/12; i++ {
			section := getSection(line, i)
			ocrLine += recognizeDigit(section)
		}
		out = append(out, ocrLine)
	}
	return out
}

func getSection(line string, index int) (out string) {
	length := len(line)
	row := length / 4
	for i := 3 * index; i < length; i += row {
		out += fmt.Sprintf("%s\n", line[i:i+3])
	}
	return out
}

func getLines(s string) (out []string) {
	linesN := countLines(s)
	lines := strings.Replace(s, "\n", "", -1)
	row := len(lines) / linesN
	for i := 0; i < linesN; i++ {
		out = append(out, lines[i*row:(i+1)*row])
	}
	return out
}

func countLines(s string) (count int) {
	for _, r := range s {
		if r == '\n' {
			count++
		}
	}
	return count / 4
}
