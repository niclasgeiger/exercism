package minesweeper

import (
	"bytes"
	"errors"
	"strconv"
)

const testVersion = 1

type MineSweeper struct {
}

func (b *Board) Count() error {
	if (*b)[0][0] != '+' {
		return errors.New("Wrong Format")
	}
	rowLen := len((*b)[0])
	for rowIndex, row := range *b {
		if len(row) != rowLen {
			return errors.New("Wrong Length")
		}
		for fieldIndex, field := range row {
			if fieldIndex == 0 && (field != '|' && field != '+') {
				return errors.New("no edge")
			}
			if fieldIndex == len(row)-1 && (field != '|' && field != '+') {
				return errors.New("no edge")
			}
			if field == ' ' {
				count := 0
				if fieldIndex > 0 && rowIndex > 0 && (*b)[rowIndex-1][fieldIndex-1] == '*' {
					count++
				}
				if fieldIndex > 0 && (*b)[rowIndex][fieldIndex-1] == '*' {
					count++
				}
				if fieldIndex > 0 && rowIndex < len((*b))-1 && (*b)[rowIndex+1][fieldIndex-1] == '*' {
					count++
				}
				if rowIndex > 0 && (*b)[rowIndex-1][fieldIndex] == '*' {
					count++
				}
				if rowIndex < len((*b))-1 && (*b)[rowIndex+1][fieldIndex] == '*' {
					count++
				}
				if fieldIndex < len(row)-1 && rowIndex > 0 && (*b)[rowIndex-1][fieldIndex+1] == '*' {
					count++
				}
				if fieldIndex < len(row)-1 && (*b)[rowIndex][fieldIndex+1] == '*' {
					count++
				}
				if fieldIndex < len(row)-1 && rowIndex < len(*b)-1 && (*b)[rowIndex+1][fieldIndex+1] == '*' {
					count++
				}
				if count > 0 {
					buffer := bytes.NewBufferString(strconv.Itoa(count))
					(*b)[rowIndex][fieldIndex] = buffer.Bytes()[0]
				} else {
					(*b)[rowIndex][fieldIndex] = ' '
				}

			} else {
				if field != '+' && field != '-' && field != '|' && field != '*' {
					return errors.New("Wrong Character")
				}
			}
		}
	}
	return nil
}
