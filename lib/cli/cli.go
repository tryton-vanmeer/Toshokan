package cli

import (
	"fmt"
	"os"
	"strings"

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
	search_terms := strings.Join(args, " ")
	fmt.Println(search_terms)
}

func Run() error {
	var rootCmd = &cobra.Command{
		Use:     "toshokan [flags] SEARCH...",
		Short:   "Search installed Steam games to find their APPID.",
		Long:    "Toshokan is a CLI tool for searching your installed Steam games to find their APPID.",
		PreRunE: preRunHelp,
		Run:     searchCmd,
	}

	rootCmd.Execute()

	return nil
}
