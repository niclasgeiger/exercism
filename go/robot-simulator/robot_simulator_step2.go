package robot

func StartRobot(commands chan Command, actions chan Action) {
	defer close(actions)
	for command := range commands {
		switch command {
		case 'A':
			actions <- ADVANCE
		case 'L':
			actions <- LEFT
		case 'R':
			actions <- RIGHT
		case ' ':
			actions <- INIT
		}
	}
}

func Room(rect Rect, robot Step2Robot, actions chan Action, robots chan Step2Robot) {
	defer close(robots)
	for action := range actions {
		robot = robot.DoAction(rect, action)
	}
	robots <- robot
}

func (robot Step2Robot) DoAction(rect Rect, action Action) Step2Robot {
	switch action {
	case INIT:
		return NewStep2Robot()
	case ADVANCE:
		robot.Advance(rect)
	case LEFT:
		robot.Left(rect)
	case RIGHT:
		robot.Right(rect)
	}
	return robot
}

func NewStep2Robot() Step2Robot {
	return Step2Robot{
		Pos: Pos{
			Northing: 1,
			Easting:  1,
		},
	}
}

func (robot *Step2Robot) Right(rect Rect) {
	switch robot.Dir {
	case N:
		robot.Dir = E
	case S:
		robot.Dir = W
	case W:
		robot.Dir = N
	case E:
		robot.Dir = S
	}
}

func (robot *Step2Robot) Left(rect Rect) {
	switch robot.Dir {
	case N:
		robot.Dir = W
	case S:
		robot.Dir = E
	case W:
		robot.Dir = S
	case E:
		robot.Dir = N
	}
}

func (robot *Step2Robot) Advance(rect Rect) {
	switch robot.Dir {
	case N:
		if robot.Northing < rect.Max.Northing {
			robot.Northing++
		}
	case S:
		if robot.Northing > rect.Min.Northing {
			robot.Northing--
		}
	case W:
		if robot.Easting > rect.Min.Easting {
			robot.Easting--
		}
	case E:
		if robot.Easting < rect.Max.Easting {
			robot.Easting++
		}
	}
}
