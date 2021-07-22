package steam

import (
	"fmt"
	"log"
	"os"

	"github.com/Jleagle/steam-go/steamvdf"
)

var STEAM_APPS_ROOT = ".steam/steam/steamapps"
var STEAM_LIBRARY_FOLDERS_VDF = "libraryfolders.vdf"

func GetLibraryFolders() []string {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(home)

	kv, err := steamvdf.ReadFile(STEAM_LIBRARY_FOLDERS_VDF)

	if err != nil {
		log.Fatalf("cannot read: %s", STEAM_LIBRARY_FOLDERS_VDF)
	}

	fmt.Println(kv)

	return []string{kv.Key}
}
