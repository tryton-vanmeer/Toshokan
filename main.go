package main

import (
	"fmt"
	"toshokan/pkg/steam"

	"github.com/manifoldco/promptui"
)

func main() {
	games := steam.GetApps()
	games.Sort()

	prompt := promptui.Select{
		Label: "Steam Games",
		Items: games,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(games[i])
}
