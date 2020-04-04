package vogue

import "time"

type eventNow struct{}

func (e eventNow) When() time.Time {
	return time.Now()
}

type eventQuit struct {
	eventNow
}
