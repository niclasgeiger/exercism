package matrix

import (
	"sort"
)

const testVersion = 1

type Pair [2]int

func (m Matrix) Saddle() (pair []Pair) {
	cols := m.Cols()
	rows := m.Rows()
	biggestRow := []int{}
	smallestCol := []int{}
	for _, row := range rows {
		sort.Sort(sort.IntSlice(row))
		biggestRow = append(biggestRow, row[len(row)-1])
	}
	for _, col := range cols {
		sort.Sort(sort.IntSlice(col))
		smallestCol = append(smallestCol, col[0])
	}
	for i := 0; i < len(m.Rows()); i++ {
		for j := 0; j < len(m.Cols()); j++ {
			row := m.Rows()[i][j]
			if row == biggestRow[i] && row == smallestCol[j] {
				pair = append(pair, Pair{i, j})
			}
		}
	}

	return pair
}
