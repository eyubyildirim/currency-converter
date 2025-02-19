package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type CurrencyModel struct {
	choices  []string
	cursor   int
	selected string
}

func (m CurrencyModel) String() string {
	return fmt.Sprintf("CurrencyModel: %s", m.selected)
}

func (m CurrencyModel) Init() tea.Cmd {
	return nil
}

func (m CurrencyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}
		case "enter", " ":
			m.selected = m.choices[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m CurrencyModel) View() string {
	s := "Select the target currency: \n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nPress q to quit.\n"

	return s
}
