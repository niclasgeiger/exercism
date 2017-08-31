package wordy

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const testVersion = 1

var operations = map[string]func(a, b int) int{
	"plus":          func(a, b int) int { return a + b },
	"minus":         func(a, b int) int { return a - b },
	"divided by":    func(a, b int) int { return a / b },
	"multiplied by": func(a, b int) int { return a * b },
}

var available = availableOperations()

func Answer(question string) (number int, ok bool) {
	if question[:8] != "What is " {
		return 0, false
	}
	question = question[8 : len(question)-1]
	ops := split(question)
	if len(ops) < 3 || len(ops)%2 == 0 {
		return 0, false
	}
	number, err := strconv.Atoi(ops[0])
	if err != nil {
		return 0, false
	}
	for i := 1; i < len(ops); i += 2 {
		b, err := strconv.Atoi(ops[i+1])
		if err != nil {
			return 0, false
		}
		if !strings.Contains(available, ops[i]) {
			return 0, false
		}
		number = operations[ops[i]](number, b)
	}
	return number, true

}

func availableOperations() (available string) {
	for op, _ := range operations {
		available += fmt.Sprintf("(%s)", op)
	}
	return available
}

func split(question string) (out []string) {
	number := true
	current := ""
	for _, r := range question {
		if unicode.IsNumber(r) || r == '-' {
			if number {
				current += string(r)
			} else {
				number = true
				current = strings.TrimSpace(current)
				out = append(out, current)
				current = string(r)
			}
		}
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			if !number {
				current += string(r)
			} else {
				number = false
				current = strings.TrimSpace(current)
				out = append(out, current)
				current = ""
			}
		}
	}
	out = append(out, current)
	return out
}
