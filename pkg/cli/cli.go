package cli

import (
	"encoding/json"
	"fmt"
	"strings"
	"toshokan/pkg/steam"

	"github.com/spf13/cobra"
)

type GameInfo struct {
	Name             string `json:"name"`
	AppID            string `json:"appid"`
	InstallDirectory string `json:"install_directory"`
	ProtonPrefix     string `json:"proton_prefix"`
}

var jsonFlag bool

var listCmd = &cobra.Command{
	Use:   "list",
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
	Use:   "search [flags] [string to search]",
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
	Use:   "info [flags] [appid]",
	Args:  cobra.ExactArgs(1),
	Short: "Get info about an installed Steam game.",
	Run: func(cmd *cobra.Command, args []string) {
		game, err := steam.GetGame(args[0])

		if err != nil {
			fmt.Println(err)
			return
		}

		if jsonFlag {
			protonPrefix := ""

			if game.IsProton() {
				protonPrefix = game.ProtonPrefix()
			}

			info := GameInfo{
				Name:             game.Name,
				AppID:            game.AppID,
				InstallDirectory: game.InstallDirectory,
				ProtonPrefix:     protonPrefix,
			}

			infoJson, _ := json.Marshal(info)

			fmt.Println(string(infoJson))

			return
		}

		builder := strings.Builder{}

		builder.WriteString(fmt.Sprintf("🎮 %s\n", game.Name))
		builder.WriteString(fmt.Sprintf("🌐 %s\n", game.GetStorePage()))

		builder.WriteString(
			fmt.Sprintf("📂 file://%s\n", game.LibraryFolder))

		if game.IsProton() {
			builder.WriteString(
				fmt.Sprintf("⚛️  file://%s", game.ProtonPrefix()))
		}

		fmt.Println(builder.String())
	},
}

func Run() {
	var rootCmd = &cobra.Command{
		Use:  "toshokan [command]",
		Long: "Toshokan is a CLI tool for interacting with your Steam library on Linux",
	}

	infoCmd.Flags().BoolVar(&jsonFlag, "json", false, "print in JSON format")

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(infoCmd)

	rootCmd.Execute()
}
