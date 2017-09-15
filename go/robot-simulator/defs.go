package robot

import "fmt"

// definitions used in step 1

const (
	N Dir = iota
	E
	S
	W
)

type Action rune

const (
	ADVANCE Action = 'A'
	LEFT    Action = 'L'
	RIGHT   Action = 'R'
	INIT    Action = ' '
	REPORT  Action = '_'
)

var Step1Robot struct {
	X, Y int
	Dir
}

type Dir int

func (d Dir) String() string {
	switch d {
	case E:
		return "East"
	case N:
		return "North"
	case S:
		return "South"
	case W:
		return "West"
	}
	return "unknown"
}

var _ fmt.Stringer = Dir(1729)

// additional definitions used in step 2

type Command byte // valid values are 'R', 'L', 'A'
type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Step2Robot struct {
	Dir
	Pos
}

// additional definition used in step 3

type Step3Robot struct {
	Name string
	Step2Robot
}

type Action3 struct {
	RobotName string
	Action
}
