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
		Use:     "toshokan",
		PreRunE: preRunHelp,
		Run:     searchCmd,
	}

	rootCmd.Execute()

	return nil
}
