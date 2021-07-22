package steam

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Jleagle/steam-go/steamvdf"
)

var STEAM_APPS_ROOT = ".steam/steam/steamapps"

// get the users configured steam libraries
func LibraryFolders() []string {
	// get the users HOME
	home, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	// create the 'libraryfolders.vdf' path
	vdf_path := fmt.Sprintf("%s/%s/%s", home, STEAM_APPS_ROOT, "libraryfolders.vdf")

	kv, err := steamvdf.ReadFile(vdf_path)

	if err != nil {
		log.Fatalf("cannot read: %s", vdf_path)
	}

	// create arrary to append libraries to, with default library path as first element
	folders := []string{
		fmt.Sprintf("%s/%s", home, STEAM_APPS_ROOT),
	}

	// if the key is an integer, it maps to a steam library
	for key := range kv.GetChildrenAsMap() {
		_, err := strconv.Atoi(key)

		if err == nil {
			child, _ := kv.GetChild(key)

			folders = append(folders, child.Value)
		}
	}

	return folders
}
