package beer

import (
	"errors"
	"fmt"
)

const testVersion = 1

const (
	plural = "%d bottles of beer on the wall, %d bottles of beer.\n" +
		"Take one down and pass it around, %d bottles of beer on the wall.\n"
	two = "2 bottles of beer on the wall, 2 bottles of beer.\n" +
		"Take one down and pass it around, 1 bottle of beer on the wall.\n"
	one = "1 bottle of beer on the wall, 1 bottle of beer.\n" +
		"Take it down and pass it around, no more bottles of beer on the wall.\n"
	none = "No more bottles of beer on the wall, no more bottles of beer.\n" +
		"Go to the store and buy some more, 99 bottles of beer on the wall.\n"
)

var (
	ErrBottleCount               = errors.New("bottles need to be in between 0 and 99")
	ErrLowerBoundBiggerThanUpper = errors.New("lower bound needs to be smaller than upper")
)

func Verse(n int) (string, error) {
	switch n {
	case 0:
		return none, nil
	case 1:
		return one, nil
	case 2:
		return two, nil
	}
	if n < 100 && n > 2 {
		return fmt.Sprintf(plural, n, n, n-1), nil
	}
	return "", ErrBottleCount
}

func Verses(upper, lower int) (string, error) {
	out := ""
	if upper < lower {
		return "", ErrLowerBoundBiggerThanUpper
	}
	for i := upper; i >= lower; i-- {
		verse, err := Verse(i)
		if err != nil {
			return "", err
		}
		out += fmt.Sprintf("%s\n", verse)
	}
	return out, nil
}

func Song() string {
	song, _ := Verses(99, 0)
	return song
}
