package lock

import "time"

type EventHandler func(e Event)

type Event struct {
	Time   time.Time
	Locked bool
}
