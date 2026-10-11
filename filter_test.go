package main

import "testing"

func TestIsMine(t *testing.T) {
	cases := []struct {
		name string
		want bool
	}{
		{"E8-A4-T2 (FM) is due", true},
		{"E8-B2-T1 (AM) is due", false},
		{"E7-D2-T1(RRS) is due", false},
		{"A1-Application-9A-Cloud-Server-Load-Balancer-Cluster-Merging-AVN is due", false},
		{"A1W9T2 is due", false},
		{"Assignment 8 is due", true},
	}

	for _, c := range cases {
		got := isMine(c.name)
		if got != c.want {
			t.Errorf("isMine(%q) = %v, want %v", c.name, got, c.want)
		}
	}
}
