package dna

import (
	"errors"
	"strings"
)

const testVersion = 2

var (
	Aminos          = "ATCG"
	ErrNoAmino      = errors.New("This is no amino acid.")
	ErrInvalidAmino = errors.New("DNA contains invalid amino acid.")
)

type DNA string

type Histogram map[rune]int

func (d DNA) Count(in byte) (count int, err error) {
	amino := rune(in)
	if !validAmino(amino) {
		return 0, ErrNoAmino
	}
	for _, r := range d {
		if !validAmino(r) {
			return 0, ErrInvalidAmino
		}
		if r == amino {
			count++
		}
	}
	return count, nil
}

func (d DNA) Counts() (histogram Histogram, err error) {
	histogram = Histogram{}
	for _, r := range Aminos {
		histogram[r], err = d.Count(byte(r))
		if err != nil {
			return histogram, err
		}
	}
	return histogram, nil
}

func validAmino(r rune) bool {
	return strings.Contains(Aminos, string(r))
}
