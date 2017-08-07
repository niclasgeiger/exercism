package clock
import (
	"fmt"
)
const testVersion = 4



type Clock struct {
	hour int
	minute int
}

func New(hour, minute int) Clock {
	for minute >= 60 {
		hour++
		minute-=60
	}
	for minute < 0 {
		minute +=60
		hour--
	}
	for hour < 0 {
		hour += 24
	}
	c  := Clock{hour%24, minute}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d",c.hour,c.minute)
}

func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	for c.minute >= 60 {
		c.hour++
		c.minute-=60
	}
	for c.minute < 0 {
		c.minute += 60
		c.hour--
	}
	for c.hour < 0 {
		c.hour += 24
	}
	c.hour %= 24
	return c
}
