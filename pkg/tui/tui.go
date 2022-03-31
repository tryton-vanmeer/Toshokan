package tui

import (
	"strings"
	"toshokan/pkg/steam"

	"github.com/manifoldco/promptui"
)

func Run() {
	games := steam.GetApps()
	games.Sort()

	details := `
{{ "App ID:" | faint }} {{ .AppID }}
{{ "Store Page:" | faint }} {{ .GetStorePage }}
{{ "Install Directory:" | faint }} {{ .InstallDirectory }}
{{ if .ProtonPrefix }}{{ "Proton Prefix:" | faint }} {{ .ProtonPrefix }}{{ end }}`

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
