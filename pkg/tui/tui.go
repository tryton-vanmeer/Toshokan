package tui

import (
	_ "embed"
	"strings"
	"toshokan/pkg/steam"

	"github.com/manifoldco/promptui"
)

//go:embed "details.tmpl"
var details string

func Run() {
	games := steam.GetApps()
	games.Sort()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "{{ .Name | underline }}",
		Inactive: "{{ .Name }}",
		Details:  details,
	}

	searcher := func(input string, index int) bool {
		game := games[index]
		name := strings.ToLower(game.Name)
		input = strings.ToLower(input)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:        "Games",
		HideSelected: true,
		Items:        games,
		Templates:    templates,
		Searcher:     searcher,
	}

	prompt.Run()
}
