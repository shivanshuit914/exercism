package clock

import "fmt"

const testVersion = 4

// You can find more details and hints in the test file.

type Clock struct {
	minute int
}

func New(hour, minute int) Clock {
	c := Clock{(60*hour + minute) % 1440}

	if c.minute < 0 {
		c.minute += 1440
	}

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/60, c.minute%60)
}

func (c Clock) Add(minutes int) Clock {
	c.minute = (c.minute + minutes) % 1440
	if c.minute < 0 {
		c.minute += 1440
	}
	return c
}
