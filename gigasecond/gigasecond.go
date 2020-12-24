
package gigasecond


import (
	"time"
)

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	time := t.Add(time.Second * time.Duration(1000000000))

	return time
}

