package main

import (
	"fmt"
	"strings"
	"toshokan/pkg/steam"

	"github.com/manifoldco/promptui"
)

func main() {
	games := steam.GetApps()
	games.Sort()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "{{ .Name | underline }}",
		Inactive: "{{ .Name }}",
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

	i, _, err := prompt.Run()

	if err != nil {
		return
	}

	fmt.Println(games[i])
}
