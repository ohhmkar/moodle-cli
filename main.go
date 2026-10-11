package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
)

func main() {
	if _, err := tea.NewProgram(model{loading: true}).Run(); err != nil {
		log.Fatal(err)
	}
}
