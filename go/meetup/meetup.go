package meetup

import (
	"sort"
	"time"
)

const testVersion = 3

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(schedule WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return getWeekday(date, weekday, schedule)
}

func getWeekday(date time.Time, weekday time.Weekday, schedule WeekSchedule) int {
	next := getNextFourWeekDays(date, weekday, schedule)
	switch schedule {
	case First:
		{
			return next[0]
		}
	case Second:
		{
			return next[1]
		}
	case Third:
		{
			return next[2]
		}
	case Fourth:
		{
			return next[3]
		}
	case Last:
		{
			sort.Sort(sort.IntSlice(next))
			return next[len(next)-1]
		}
	case Teenth:
		{
			for _, day := range next {
				if day > 12 && day < 20 {
					return day
				}
			}
		}
	}
	return -1
}

func getNextFourWeekDays(date time.Time, weekday time.Weekday, schedule WeekSchedule) (out []int) {
	for date.Weekday() != weekday {
		date = date.AddDate(0, 0, 1)
	}
	for i := 0; i < 5; i++ {
		out = append(out, date.Day())
		date = date.AddDate(0, 0, 7)
	}
	return out
}
