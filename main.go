package main

import (
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "login" {
		if err := login(); err != nil {
			log.Fatal(err)
		}
		return
	}
	if _, err := tea.NewProgram(model{loading: true}).Run(); err != nil {
		log.Fatal(err)
	}
}
