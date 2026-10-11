package main

import (
	"testing"
	"time"
)

func TestHumanize(t *testing.T) {
	cases := []struct {
		d    time.Duration
		want string
	}{
		{-2 * time.Hour, "overdue"},
		{45 * time.Minute, "in 45 m"},
		{9 * time.Hour, "in 9 h"},
		{50 * time.Hour, "in 2 d"},
	}

	for _, c := range cases {
		got := humanize(c.d)
		if got != c.want {
			t.Errorf("humanize(%v) = %q, want %q", c.d, got, c.want)
		}
	}
}
