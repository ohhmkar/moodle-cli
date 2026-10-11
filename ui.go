package main

import (
	"fmt"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	urgentStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5f87")) // red for abhi karle bhai
	soonStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffd75f")) // yellow for kinda close
	laterStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#5fd787")) // green for time hai bohot
)

type model struct {
	events []Event
	cursor int //which row is selected
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
			return m, nil
		case "down", "j":
			if m.cursor < len(m.events)-1 {
				m.cursor++
			}
			return m, nil
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	var b strings.Builder
	b.WriteString("Your deadlines\n\n")

	for i, e := range m.events {
		prefix := "  "
		if i == m.cursor {
			prefix = "> "
		}
		due := time.Unix(e.TimeSort, 0)
		left := time.Until(due)
		fmt.Fprintf(&b, "%s%-30s%-30s %s\n", prefix, e.Name, e.Course.Shortname, urgencyStyle(left).Render(humanize(left)))
	}

	b.WriteString("\n↑/↓ move . q quit\n")
	return tea.NewView(b.String())
}
