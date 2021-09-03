package steam

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"toshokan/src/util"

	"github.com/Jleagle/steam-go/steamvdf"
)

var STEAM_APPS_ROOT = ".steam/steam/steamapps"

type App struct {
	Name             string
	AppID            string
	LibraryFolder    string
	InstallDirectory string
}

func (app App) ToString() string {
	builder := strings.Builder{}

	builder.WriteString(app.Name)
	builder.WriteString(" ")
	builder.WriteString(app.AppID)
	builder.WriteString(" [")
	builder.WriteString(
		util.FileHyperlink(app.InstallDirectory, "Install Directory"))
	builder.WriteString("]")

	if app.isProton() {
		builder.WriteString(" [")
		builder.WriteString(
			util.FileHyperlink(app.protonPrefix(), "Proton Prefix"))
		builder.WriteString("]")
	}

	return builder.String()
}

// get path game would use for it's proton prefix
func (app App) protonPrefix() string {
	return fmt.Sprintf("%s/compatdata/%s", app.LibraryFolder, app.AppID)
}

// check if game uses proton
func (app App) isProton() bool {
	// if the proton prefix path exists, assume it is using proton
	// game can still be using proton, but hasn't performed first launch (so path won't exist)
	_, err := os.Stat(app.protonPrefix())

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

	// add the default library
	directories = append(directories,
		fmt.Sprintf("%s/%s", home, STEAM_APPS_ROOT))

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
func installedGames() (games []App) {
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
				games = append(games, game)
			}
		}
	}

	return
}

func SearchInstalledGames(search string) (games []App) {
	installed_games := installedGames()

	for _, game := range installed_games {
		if util.StringContains(game.Name, search) {
			games = append(games, game)
		}
	}

	return
}
