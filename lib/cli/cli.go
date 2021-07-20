package cli

import (
	"log"

	"github.com/tryton-vanmeer/toshokan/lib/utils"
	"github.com/urfave/cli"
)

func Run(args []string) {
	log.SetFlags(0)

	app := &cli.App{
		Name:  "toshokan",
		Usage: "CLI tool for finding APPID for Steam games",
	}

	err := app.Run(args)

	if err != nil {
		utils.Error(err)
	}
}
