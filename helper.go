package main

import (
	"fmt"
	"time"

	"charm.land/lipgloss/v2"
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

func urgencyStyle(d time.Duration) lipgloss.Style {
	h := d.Hours()
	if h < 24 {
		return urgentStyle
	} else if h < 72 {
		return soonStyle
	} else {
		return laterStyle
	}
}
