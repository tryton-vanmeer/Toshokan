package cli

import (
	"fmt"
	"os"
	"strings"
	"toshokan/lib/steam"

	"github.com/spf13/cobra"
)

func preRunHelp(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}

	return nil
}

func searchCmd(cmd *cobra.Command, args []string) {
	games := steam.SearchInstalledGames(strings.Join(args, " "))

	for _, game := range games {
		fmt.Printf("%s (%s) [%s]\n", game.Name, game.AppID, game.LibraryFolder)
	}
}

func Run() error {
	var rootCmd = &cobra.Command{
		Use:     "toshokan [flags] SEARCH...",
		Short:   "Search installed Steam games to find their APPID.",
		Long:    "Toshokan is a CLI tool for searching your installed Steam games to find their APPID.",
		Example: "toshokan half life",
		PreRunE: preRunHelp,
		Run:     searchCmd,
	}

	rootCmd.Execute()

	return nil
}
