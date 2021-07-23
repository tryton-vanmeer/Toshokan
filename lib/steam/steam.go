package steam

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Jleagle/steam-go/steamvdf"
)

var STEAM_APPS_ROOT = ".steam/steam/steamapps"

type App struct {
	name          string
	appid         string
	libraryFolder string
}

// get the users configured steam libraries
func LibraryFolders() (directories []string) {
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
			directories = append(directories, child.Value+"/steamapps")
		}
	}

	return
}

// parse an appmanifest_$id.acf and return a Game object
func ParseAppManifest(libraryFolder string, filename string) App {
	app_manifest_path := fmt.Sprintf("%s/%s", libraryFolder, filename)
	kv, err := steamvdf.ReadFile(app_manifest_path)
	if err != nil {
		log.Fatalf("error reading %s", app_manifest_path)
	}

	app_manifest := kv.GetChildrenAsMap()

	return App{
		name:          app_manifest["name"],
		appid:         app_manifest["appid"],
		libraryFolder: libraryFolder,
	}
}

// return a list of the users installed apps
func InstalledGames() (games []App) {
	folders := LibraryFolders()

	// search library directories for app manifest files
	for _, folder := range folders {
		files, err := ioutil.ReadDir(folder)

		if err != nil {
			log.Fatalf("error reading directory: %s", folder)
		}

		for _, file := range files {
			if strings.Contains(file.Name(), "appmanifest") {
				game := ParseAppManifest(folder, file.Name())
				games = append(games, game)
			}
		}
	}

	return
}
