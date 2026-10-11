package main

import (
	"fmt"
	"time"
)

func humanize(d time.Duration) string {
	if d < 0 {
		return "overdue"
	}
	if d < time.Hour {
		return fmt.Sprintf("in %d m", int(d.Minutes()))
	}
	if d < 24*time.Hour {
		return fmt.Sprintf("in %d h", int(d.Hours()))
	}
	return fmt.Sprintf("in %d d", int(d.Hours()/24))
}
