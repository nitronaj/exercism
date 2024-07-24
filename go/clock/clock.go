package clock

import (
	"time"
)

// Define the Clock type here.
type Clock struct {
	m int
	h int
}

func New(h, m int) Clock {
	t := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), h, m, 0, 0, time.UTC)
	return Clock{
		h: t.Hour(),
		m: t.Minute(),
	}

}

func (c Clock) Add(m int) Clock {
	t := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), c.h, c.m, 0, 0, time.UTC)
	d := time.Minute * time.Duration(m)
	currentTime := t.Add(d)
	c.h = currentTime.Hour()
	c.m = currentTime.Minute()
	return c
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	t := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), c.h, c.m, 0, 0, time.UTC)
	time := t.Format("15:04")
	return time
}
