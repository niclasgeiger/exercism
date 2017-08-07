package queenattack

import (
	"errors"
	"strconv"
)

const testVersion = 2

var (
	conv = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
	}
)

func CanQueenAttack(queen, peasant string) (bool, error) {
	if queen == peasant {
		return false, errors.New("same field!")
	}
	q, err := parsePosition(queen)
	if err != nil {
		return false, err
	}
	p, err := parsePosition(peasant)
	if err != nil {
		return false, err
	}
	return canAttack(p, q)
}

func parsePosition(pos string) ([]int, error) {
	if len(pos) > 2 {
		return nil, errors.New("no valid position")
	}
	x, ok := conv[pos[:1]]
	y, err := strconv.Atoi(pos[1:])
	if x > 8 || x < 1 || y > 8 || y < 1 || err != nil || !ok {
		return nil, errors.New("out of the board")
	}
	return []int{x, y}, nil
}

func canAttack(queen, peasant []int) (bool, error) {
	if queen[0] == peasant[0] || queen[1] == peasant[1] {
		return true, nil
	}
	for i := -1; i < 2; i += 2 {
		for j := -1; j < 2; j += 2 {
			var temp []int = make([]int, 2)
			copy(temp, queen)
			for temp[0] > 0 && temp[0] < 9 && temp[1] > 0 && temp[1] < 9 {
				temp[0] += i
				temp[1] += j
				if temp[0] == peasant[0] && temp[1] == peasant[1] {
					return true, nil
				}
			}
		}
	}
	return false, nil
}
