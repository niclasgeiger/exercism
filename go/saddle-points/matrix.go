package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	matrix string
	cols   [][]int
	rows   [][]int
}

func New(in string) (m *Matrix, err error) {
	r, err := rows(in)
	if err != nil {
		return nil, err
	}
	c, err := cols(in)
	if err != nil {
		return nil, err
	}
	m = &Matrix{
		matrix: in,
		cols:   c,
		rows:   r,
	}
	return m, nil
}

func rows(in string) (out [][]int, err error) {
	lines := strings.Split(in, "\n")
	prevCount := len(strings.Split(lines[0], " "))
	for _, line := range lines {
		row := []int{}
		elements := splitLine(line)
		if prevCount != len(elements) {
			return nil, errors.New("different row length")
		}
		for _, s := range elements {
			val, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			row = append(row, val)
		}
		out = append(out, row)
		prevCount = len(elements)
	}
	return out, nil
}

func cols(in string) (out [][]int, err error) {
	lines := strings.Split(in, "\n")
	for range strings.Split(lines[0], " ") {
		out = append(out, []int{})
	}
	prevCount := len(strings.Split(lines[0], " "))
	for _, line := range lines {
		elements := splitLine(line)
		if prevCount != len(elements) {
			return nil, errors.New("different row length")
		}
		for i, s := range elements {
			val, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			out[i] = append(out[i], val)
		}
		prevCount = len(elements)
	}
	return out, nil
}

func (m *Matrix) Rows() (out [][]int) {
	out = [][]int{}
	for i := 0; i < len(m.rows); i++ {
		out = append(out, []int{})
	}
	for i := 0; i < len(m.rows); i++ {
		for j := 0; j < len(m.rows[0]); j++ {
			out[i] = append(out[i], m.rows[i][j])
		}
	}
	return out
}

func (m *Matrix) Cols() (out [][]int) {
	out = [][]int{}
	for i := 0; i < len(m.cols); i++ {
		out = append(out, []int{})
	}
	for i := 0; i < len(m.cols); i++ {
		for j := 0; j < len(m.cols[0]); j++ {
			out[i] = append(out[i], m.cols[i][j])
		}
	}
	return out
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m.rows) || col < 0 || col >= len(m.cols) {
		return false
	}
	m.cols[col][row] = val
	m.rows[row][col] = val
	return true
}

func splitLine(line string) []string {
	line = strings.TrimSpace(line)
	return strings.Split(line, " ")
}
