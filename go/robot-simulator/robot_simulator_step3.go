package robot

import (
	"strings"
	"sync"
)

type Robots struct {
	Map map[string]Step3Robot
	sync.Mutex
}

func Room3(rect Rect, robots []Step3Robot, actions chan Action3, report chan []Step3Robot, log chan string) {
	defer close(report)
	robotMap := Robots{
		Map: make(map[string]Step3Robot),
	}
	for _, robot := range robots {
		robotMap.Mutex.Lock()
		if _, found := robotMap.Map[robot.Name]; found {
			log <- "same named robots"
			report <- nil
			return
		}
		if len(robot.Name) == 0 {
			log <- "no name on robot"
			report <- nil
			return
		}
		if cellIsOccupied(robot.Northing, robot.Easting, robotMap.Map) {
			log <- "cell is already occupied"
			report <- nil
			return
		}
		if robot.Northing > rect.Max.Northing || robot.Northing < rect.Min.Northing || robot.Easting > rect.Max.Easting || robot.Easting < rect.Min.Easting {
			log <- "cell is already occupied"
			report <- nil
			return
		}
		robotMap.Map[robot.Name] = robot
		robotMap.Mutex.Unlock()
	}
	for action := range actions {
		robotMap.Mutex.Lock()
		if action.Action == REPORT {
			report <- toArray(robotMap.Map)
		} else {
			robot := robotMap.Map[action.RobotName]
			robot.DoAction(rect, action.Action, log, robotMap.Map)
			robotMap.Map[action.RobotName] = robot
		}
		robotMap.Mutex.Unlock()
	}

}
func toArray(robots map[string]Step3Robot) (out []Step3Robot) {
	for _, robot := range robots {
		out = append(out, robot)
	}
	return
}

func cellIsOccupied(northing, easting RU, robots map[string]Step3Robot) bool {
	for _, robot := range robots {
		if robot.Northing == northing && robot.Easting == easting {
			return true
		}
	}
	return false
}

func StartRobot3(name string, script string, actions chan Action3, log chan string) {
	for _, command := range script {
		action := Action3{
			RobotName: name,
			Action:    (Action)(command),
		}
		actions <- action
		if !strings.Contains("ARL_ ", string(command)) {
			log <- "unknown command"
			break
		}
	}
	report := Action3{
		RobotName: name,
		Action:    REPORT,
	}
	actions <- report
}

func (robot *Step3Robot) DoAction(rect Rect, action Action, log chan string, robots map[string]Step3Robot) {
	switch action {
	case INIT:
		robot = &Step3Robot{}
		return
	case ADVANCE:
		robot.Advance(rect, log, robots)
		return
	case LEFT:
		robot.Left(rect)
		return
	case RIGHT:
		robot.Right(rect)
		return
	}
}

func (robot *Step3Robot) Advance(rect Rect, log chan string, robots map[string]Step3Robot) {
	switch robot.Dir {
	case N:
		if robot.Northing < rect.Max.Northing && !cellIsOccupied(robot.Northing+1, robot.Easting, robots) {
			robot.Northing++
		} else {
			log <- "bumped into something"
		}
	case S:
		if robot.Northing > rect.Min.Northing && !cellIsOccupied(robot.Northing-1, robot.Easting, robots) {
			robot.Northing--
		} else {
			log <- "bumped into something"
		}
	case W:
		if robot.Easting > rect.Min.Easting && !cellIsOccupied(robot.Northing, robot.Easting-1, robots) {
			robot.Easting--
		} else {
			log <- "bumped into something"
		}
	case E:
		if robot.Easting < rect.Max.Easting && !cellIsOccupied(robot.Northing, robot.Easting+1, robots) {
			robot.Easting++
		} else {
			log <- "bumped into something"
		}
	}
}
