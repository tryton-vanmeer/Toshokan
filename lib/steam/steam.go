package steam

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Jleagle/steam-go/steamvdf"
)

var STEAM_APPS_ROOT = ".steam/steam/steamapps"

type Game struct {
	name  string
	appid int
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
func ParseAppManifest(path string) Game {
	return Game{}
}

// return a list of the users installed apps
func InstalledGames() {

}
