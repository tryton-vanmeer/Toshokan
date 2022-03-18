package tui

import (
	"fmt"
	"os"
	"toshokan/pkg/steam"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list     list.Model
	showInfo bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			m.showInfo = !m.showInfo
			return m, nil
		}
	}

	if m.showInfo {
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.showInfo {
		return m.list.SelectedItem().(item).Info()
	}

	return m.list.View()
}

func Run() {
	items := getItemListFromGames(steam.GetApps())

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.SetShowStatusBar(false)
	l.SetShowTitle(false)

	m := model{l, false}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
