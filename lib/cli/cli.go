package cli

import (
	"fmt"
	"os"
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
	// search_terms := strings.Join(args, " ")

	library_folders, _ := steam.LibraryFolders()

	fmt.Println(library_folders)
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
