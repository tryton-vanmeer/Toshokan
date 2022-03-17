package tui

import (
	"fmt"
	"os"
	"toshokan/pkg/steam"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item steam.App

func (i item) Title() string       { return i.Name }
func (i item) FilterValue() string { return i.Name }
func (i item) Description() string { return i.AppID }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

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
	return docStyle.Render(m.list.View())
}

func Run() {
	items := []list.Item{}
	games := steam.GetApps()

	for _, game := range games {
		item := item{
			Name:             game.Name,
			AppID:            game.AppID,
			LibraryFolder:    game.LibraryFolder,
			InstallDirectory: game.InstallDirectory,
		}

		items = append(items, item)
	}

	m := model{list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Toshokan"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
