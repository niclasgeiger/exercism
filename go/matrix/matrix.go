package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

const testVersion = 1

type Matrix struct {
	rows [][]int
	cols [][]int
}

func New(in string) (m *Matrix, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error in: %s\n", in)
		}
	}()
	m = &Matrix{}
	lines := strings.Split(in, "\n")
	m.cols = [][]int{}
	for i := 0; i < len(lines[0]); i++ {
		m.cols = append(m.cols, []int{})
		lines[i] = strings.TrimSpace(lines[i])
	}
	fmt.Printf("%s\n", lines)
	for _, line := range lines {
		row := []int{}
		for _, s := range strings.Split(line, " ") {
			fmt.Println(s)
			val, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			row = append(row, val)
		}
		m.rows = append(m.rows, row)
	}
	for i, row := range m.rows {
		for _, col := range row {
			m.cols[i] = append(m.cols[i], col)
		}
	}
	return m, nil
}

func (m Matrix) Rows() [][]int {
	return m.rows
}

func (m Matrix) Cols() [][]int {
	return m.cols
}

func (m *Matrix) Set(row, col, val int) bool {
	return true
}
