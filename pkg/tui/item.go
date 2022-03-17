package tui

import "toshokan/pkg/steam"

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
