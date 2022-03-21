package main

import (
	"fmt"
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
		Selected: "{{ .Name }}",
	}

	prompt := promptui.Select{
		Label:     "Games",
		Items:     games,
		Templates: templates,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(games[i])
}
