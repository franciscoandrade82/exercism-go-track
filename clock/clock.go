package clock

import "fmt"

// Clock represents a simple clock
type Clock struct {
	hours, minutes int
}

// New creates a new Clock
func New(h, m int) Clock {
	hours, minutes := normalizeTime(h, m)
	return Clock{hours: hours, minutes: minutes}
}

// Add adds minutes to the clock
func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes+m)
}

// Subtract subtracts minutes from the clock
func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes-m)
}

// String returns the string representation of the clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

// normalizeTime normalizes the time to ensure proper values
func normalizeTime(h, m int) (int, int) {
	totalMinutes := h*60 + m
	totalMinutes %= 1440 // 1440 minutes in a day
	if totalMinutes < 0 {
		totalMinutes += 1440
	}
	hours := totalMinutes / 60
	minutes := totalMinutes % 60
	return hours, minutes
}

// Equal checks if two clocks represent the same time
func (c Clock) Equal(other Clock) bool {
	return c.hours == other.hours && c.minutes == other.minutes
}
