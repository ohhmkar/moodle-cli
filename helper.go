package main

import (
	"fmt"
	"os/exec"
	"runtime"
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

func openURL(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileprotocolHandler", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	return cmd.Start()
}
