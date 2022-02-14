package cli

import (
	"fmt"
	"strings"
	"toshokan/pkg/steam"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [flags]",
	Args:  cobra.NoArgs,
	Short: "List installed Steam games.",
	Run: func(cmd *cobra.Command, args []string) {
		games := steam.InstalledGames()

		for _, game := range games {
			fmt.Printf("%s (%s)\n", game.Name, game.AppID)
		}
	},
}

var searchCmd = &cobra.Command{
	Use:   "search [flags] ARGS...",
	Args:  cobra.MinimumNArgs(1),
	Short: "Search installed Steam games.",
	Run: func(cmd *cobra.Command, args []string) {
		games := steam.SearchInstalledGames(strings.Join(args, " "))

		for _, game := range games {
			fmt.Printf("%s (%s)\n", game.Name, game.AppID)
		}
	},
}

var infoCmd = &cobra.Command{
	Use:   "info [flags] APPID",
	Args:  cobra.ExactArgs(1),
	Short: "Get info about an installed Steam game.",
	Run: func(cmd *cobra.Command, args []string) {
		game, err := steam.GetGame(args[0])

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s (%s)\n", game.Name, game.AppID)
	},
}

func Run() {
	var rootCmd = &cobra.Command{
		Long: "Toshokan is a CLI tool for interacting with your Steam library on Linux",
	}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(infoCmd)

	rootCmd.Execute()
}
