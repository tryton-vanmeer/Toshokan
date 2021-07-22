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

	vdf_path := fmt.Sprintf("%s/%s/%s", home, STEAM_APPS_ROOT, STEAM_LIBRARY_FOLDERS_VDF)
	kv, err := steamvdf.ReadFile(vdf_path)

	if err != nil {
		log.Fatalf("cannot read: %s", STEAM_LIBRARY_FOLDERS_VDF)
	}

	fmt.Println(kv)

	return []string{kv.Key}
}
