package tui

import (
	"toshokan/pkg/steam"

	"github.com/charmbracelet/bubbles/list"
)

type item struct {
	app steam.App
}

type itemList []item

func (i item) FilterValue() string {
	return i.app.Name
}

func (i item) Title() string {
	return i.app.Name
}

func (i item) Description() string {
	return i.app.AppID
}

func getItemListFromGames(games []steam.App) (items []list.Item) {
	for _, game := range games {
		items = append(items, item{game})
	}

	return
}
