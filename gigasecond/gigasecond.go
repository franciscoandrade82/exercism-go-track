package gigasecond

import (
	"time"
)

// gigasecond represents a duration of 1M seconds
const gigasecond = 1_000_000_000

// AddGigasecond add gigasecond to a time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(gigasecond))
}
