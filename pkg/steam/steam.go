package steam

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"toshokan/pkg/util"

	"github.com/Jleagle/steam-go/steamvdf"
)

var STEAM_APPS_ROOT = ".steam/steam/steamapps"

type App struct {
	Name             string
	AppID            string
	LibraryFolder    string
	InstallDirectory string
}

// get the URL for the apps page on Steam
func (app App) GetStorePage() string {
	return fmt.Sprintf("https://store.steampowered.com/app/%s", app.AppID)
}

// get path game would use for it's proton prefix
func (app App) ProtonPrefix() string {
	return fmt.Sprintf("%s/compatdata/%s", app.LibraryFolder, app.AppID)
}

// check if game uses proton
func (app App) IsProton() bool {
	// if the proton prefix path exists, assume it is using proton
	// game can still be using proton, but hasn't performed first launch (so path won't exist)
	_, err := os.Stat(app.ProtonPrefix())

	return err == nil
}

// get the users configured steam libraries
func libraryFolders() (directories []string) {
	// get the users HOME
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// create the 'libraryfolders.vdf' path and parse it
	vdf_path := fmt.Sprintf("%s/%s/%s", home, STEAM_APPS_ROOT, "libraryfolders.vdf")
	kv, err := steamvdf.ReadFile(vdf_path)
	if err != nil {
		log.Fatalf("error reading %s", vdf_path)
	}

	// if the key is an integer, it maps to a steam library
	for key := range kv.GetChildrenAsMap() {
		_, err := strconv.Atoi(key)

		if err == nil {
			child, _ := kv.GetChild(key)
			path, _ := child.GetChild("path")
			directories = append(directories, path.Value+"/steamapps")
		}
	}

	return
}

// parse an appmanifest_$id.acf and return a Game object
func parseAppManifest(libraryFolder string, filename string) App {
	app_manifest_path := fmt.Sprintf("%s/%s", libraryFolder, filename)
	kv, err := steamvdf.ReadFile(app_manifest_path)
	if err != nil {
		log.Fatalf("error reading %s", app_manifest_path)
	}

	app_manifest := kv.GetChildrenAsMap()

	return App{
		Name:             app_manifest["name"],
		AppID:            app_manifest["appid"],
		LibraryFolder:    libraryFolder,
		InstallDirectory: strings.Join([]string{libraryFolder, "common", app_manifest["installdir"]}, "/"),
	}
}

// return a list of the users installed apps
func InstalledGames() (games map[string]App) {
	games = make(map[string]App)

	folders := libraryFolders()

	// search library directories for app manifest files
	for _, folder := range folders {
		files, err := ioutil.ReadDir(folder)

		if err != nil {
			log.Fatalf("error reading directory: %s", folder)
		}

		for _, file := range files {
			if strings.Contains(file.Name(), "appmanifest") {
				game := parseAppManifest(folder, file.Name())
				games[game.AppID] = game
			}
		}
	}

	return
}

func SearchInstalledGames(search string) (games []App) {
	installed_games := InstalledGames()

	for _, game := range installed_games {
		if util.StringContains(game.Name, search) {
			games = append(games, game)
		}
	}

	return
}

func GetGame(appid string) (App, error) {
	installed_games := InstalledGames()
	game, exists := installed_games[appid]

	if !exists {
		return App{}, fmt.Errorf("no game found with appid: %s", appid)
	}

	return game, nil
}
