// Package clause.
package gigasecond
// import clause.
import (
	"time"
)

// Constant declaration.
const testVersion = 4

// API function.  It uses a type from the Go standard library.
func AddGigasecond(t time.Time) time.Time {
	duration, _ := time.ParseDuration("1000000000s")
	var future time.Time = t.Add(duration)
	return future
}

