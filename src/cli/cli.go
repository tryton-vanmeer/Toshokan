package cli

import (
	"fmt"
	"strings"
	"toshokan/src/steam"
	"toshokan/src/util"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search [flags] ARGS...",
	Args:    cobra.MinimumNArgs(1),
	Short:   "Search installed Steam games to find their APPID",
	Example: "toshokan half life",
	Run: func(cmd *cobra.Command, args []string) {
		games := steam.SearchInstalledGames(strings.Join(args, " "))

		for _, game := range games {
			fmt.Printf("%s (%s) %s\n",
				game.Name,
				game.AppID,
				util.FileHyperlink(game.InstallDirectory, "Install Directory"),
			)
		}
	},
}

func Run() {
	var rootCmd = &cobra.Command{
		Long: "Toshokan is a CLI tool for interacting with your Steam library on Linux",
	}

	rootCmd.AddCommand(searchCmd)

	rootCmd.Execute()
}
