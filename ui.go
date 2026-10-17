package main

import (
	"fmt"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	urgentStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5f87"))                                                                // red for abhi karle bhai
	soonStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffd75f"))                                                                // yellow for kinda close
	laterStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#5fd787"))                                                                // green for time hai bohot
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ffffff")).Background(lipgloss.Color("#7d56f4")).Padding(0, 1) // purple for the title
	selectedStyle = lipgloss.NewStyle().Background(lipgloss.Color("#3a3a5a")).Bold(true)                                                     // highlight for selected assignment
	boxStyle      = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2)                                                       // outerbox
	helpStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#6c6c6c"))                                                                // idk bro
)

type model struct {
	events  []Event
	tab     int // 0 = upcoming 1 = overdue
	cursor  int //which row is selected
	loading bool
	err     error
}

type deadlineMsg struct {
	events []Event
	err    error
}

func (m model) Init() tea.Cmd {
	return func() tea.Msg {
		e, err := upcomingDeadlines()
		return deadlineMsg{events: e, err: err}
	}
}

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
			if m.cursor < len(m.visible())-1 {
				m.cursor++
			}
			return m, nil
		case "enter", "o":
			v := m.visible()
			if len(v) == 0 {
				return m, nil
			}
			url := v[m.cursor].URL
			openURL(url)
			return m, nil
		case "tab", "t":
			m.tab = 1 - m.tab
			m.cursor = 0
		case "r":
			m.loading = true
			return m, m.Init()
		}
	case deadlineMsg:
		m.loading = false
		m.events = msg.events
		m.err = msg.err
		m.cursor = 0
	}
	return m, nil
}

func (m model) View() tea.View {
	var b strings.Builder
	b.WriteString(titleStyle.Render("Omkar's SPIT Moodle TUI"))
	labels := []string{"Upcoming", "Overdue"}
	for i, l := range labels {
		if i == m.tab {
			b.WriteString(selectedStyle.Render(" " + l + " "))
		} else {
			b.WriteString(helpStyle.Render(" " + l + " "))
		}
	}
	b.WriteString("\n\n")

	if m.loading {
		b.WriteString("Loading... ")
		return tea.NewView(boxStyle.Render(b.String()))
	}
	if m.err != nil {
		b.WriteString(m.err.Error())
		return tea.NewView(boxStyle.Render(b.String()))
	}

	if len(m.visible()) == 0 {
		b.WriteString("nothing here!")
		return tea.NewView(boxStyle.Render(b.String()))
	}

	for i, e := range m.visible() {
		prefix := "  "
		if i == m.cursor {
			prefix = "> "
		}
		due := time.Unix(e.TimeSort, 0)
		left := time.Until(due)
		row := fmt.Sprintf("%s%-30s %-32s %-17s", prefix, e.Name, e.Course.Shortname, due.Format("Mon 02 Jan 15:04"))
		if i == m.cursor {
			row = selectedStyle.Render(row)
		}
		// countdown added after the highlight so the two styles don't clash
		fmt.Fprintf(&b, "%s %s\n", row, urgencyStyle(left).Render(humanize(left)))
	}

	b.WriteString("\n")
	b.WriteString(helpStyle.Render("↑/↓ move • q quit • <enter> open • <tab>|t tab • r refresh"))
	return tea.NewView(boxStyle.Render(b.String()))
}
