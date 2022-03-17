package tui

import (
	"sort"
	"toshokan/pkg/steam"

	"github.com/charmbracelet/bubbles/list"
)

type item struct {
	app             steam.App
	showDescription bool
}

type itemList []list.Item

func newItem(game steam.App) item {
	return item{
		app:             game,
		showDescription: false,
	}
}

func (i item) FilterValue() string {
	return i.app.Name
}

func (i item) Title() string {
	return i.app.Name
}

func (i item) Description() string {
	return i.app.AppID
}

func getItemListFromGames(games []steam.App) (items itemList) {
	for _, game := range games {
		items = append(items, newItem(game))
	}

	sort.Sort(itemList(items))

	return
}

func (items itemList) Len() int {
	return len(items)
}

func (items itemList) Less(i, j int) bool {
	return items[i].(item).app.Name < items[j].(item).app.Name
}

func (items itemList) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}
