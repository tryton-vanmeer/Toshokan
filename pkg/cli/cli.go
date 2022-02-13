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
			fmt.Println(game.ToString())
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
			fmt.Println(game.ToString())
		}
	},
}

func Run() {
	var rootCmd = &cobra.Command{
		Long: "Toshokan is a CLI tool for interacting with your Steam library on Linux",
	}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(searchCmd)

	rootCmd.Execute()
}
