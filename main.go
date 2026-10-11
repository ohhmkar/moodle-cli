package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
)

func main() {
	events, err := upcomingDeadlines()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tea.NewProgram(model{events: events}).Run(); err != nil {
		log.Fatal(err)
	}
}
