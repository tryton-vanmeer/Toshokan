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
func GetLibraryFolders() []string {
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

	folders := []string{}

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
