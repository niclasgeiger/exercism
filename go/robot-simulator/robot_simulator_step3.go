package robot

func Room3(rect Rect, robots []Step3Robot, actions chan Action3, report chan []Step3Robot, log chan string) {
	defer close(report)
	robotMap := map[string]Step3Robot{}
	for _, robot := range robots {
		robotMap[robot.Name] = robot
	}
	action := <-actions
	if len(action.RobotName) > 0 {
		robot := robotMap[action.RobotName]
		robot.DoAction(rect, action.Action, log)
		robotMap[action.RobotName] = robot
	} else {
		log <- "empty name"
	}
	for action := range actions {
		if len(action.RobotName) > 0 {
			if action.Action == REPORT {
				report <- toArray(robotMap)
			} else {
				robot := robotMap[action.RobotName]
				robot.DoAction(rect, action.Action, log)
				robotMap[action.RobotName] = robot
			}
		} else {
			log <- "empty name"
		}
	}
}
func toArray(robots map[string]Step3Robot) (out []Step3Robot) {
	for _, robot := range robots {
		out = append(out, robot)
	}
	return
}

func StartRobot3(name string, script string, actions chan Action3, log chan string) {
	for _, command := range script {
		action := Action3{
			RobotName: name,
			Action:    (Action)(command),
		}
		actions <- action
	}
	reportAction := Action3{
		RobotName: name,
		Action:    REPORT,
	}
	actions <- reportAction
}

func (robot *Step3Robot) DoAction(rect Rect, action Action, log chan string) {
	switch action {
	case INIT:
		robot = &Step3Robot{}
		return
	case ADVANCE:
		robot.Advance(rect, log)
		return
	case LEFT:
		robot.Left(rect)
		return
	case RIGHT:
		robot.Right(rect)
		return
	}
	log <- "unknown command"
}

func (robot *Step3Robot) Advance(rect Rect, log chan string) {
	switch robot.Dir {
	case N:
		if robot.Northing < rect.Max.Northing {
			robot.Northing++
		} else {
			log <- "bumped into wall"
		}
	case S:
		if robot.Northing > rect.Min.Northing {
			robot.Northing--
		} else {
			log <- "bumped into wall"
		}
	case W:
		if robot.Easting > rect.Min.Easting {
			robot.Easting--
		} else {
			log <- "bumped into wall"
		}
	case E:
		if robot.Easting < rect.Max.Easting {
			robot.Easting++
		} else {
			log <- "bumped into wall"
		}
	}
}
