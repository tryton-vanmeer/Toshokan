package tui

import (
	"toshokan/pkg/steam"

	"github.com/charmbracelet/bubbles/list"
)

type Item struct {
	app steam.App
}

func (i Item) FilterValue() string {
	return i.app.Name
}

func (i Item) Title() string {
	return i.app.Name
}

func (i Item) Description() string {
	return i.app.AppID
}

func GetItemListFromGames(games []steam.App) (items []list.Item) {
	for _, game := range games {
		items = append(items, Item{game})
	}

	return
}
