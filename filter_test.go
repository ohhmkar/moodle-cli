package main

import "testing"

func TestIsMine(t *testing.T) {
	t.Setenv("MYBATCH", "A4")
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
		{"Batch A4 is due", true},
		{"Batch B4 is due", false},
		{"E5C3T1 is due", false},
		{"E6T1C3 is due", false},
		{"W6A4T2 is due", true},
		{"E5C4 T2 is due", false},
	}

	for _, c := range cases {
		got := isMine(c.name)
		if got != c.want {
			t.Errorf("isMine(%q) = %v, want %v", c.name, got, c.want)
		}
	}
}
