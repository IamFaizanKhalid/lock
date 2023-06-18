package lock

import "time"

type Event struct {
	Time   time.Time
	Locked bool
}
