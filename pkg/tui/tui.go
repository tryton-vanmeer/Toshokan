package tui

import (
	"fmt"
	"os"
	"toshokan/pkg/steam"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list list.Model
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
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}

func Run() {
	items := getItemListFromGames(steam.GetApps())
	delegate := list.NewDefaultDelegate()
	// delegate.ShowDescription = false

	l := list.New(items, delegate, 0, 0)
	l.SetShowStatusBar(false)
	l.SetShowTitle(false)

	m := model{l}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
