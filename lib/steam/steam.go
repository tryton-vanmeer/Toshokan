package steam

import (
	"fmt"

	"github.com/Jleagle/steam-go/steamvdf"
)

var STEAM_APPS_ROOT = ".steam/steam/steamapps"
var STEAM_LIBRARY_FOLDERS_VDF = fmt.Sprintf("%s/libraryfolders.vdf", STEAM_APPS_ROOT)

func GetLibraryFolders() ([]string, error) {
	kv, err := steamvdf.ReadFile(STEAM_LIBRARY_FOLDERS_VDF)

	if err != nil {
		return nil, fmt.Errorf("unable to read: %s", STEAM_LIBRARY_FOLDERS_VDF)
	}

	fmt.Println(kv)

	return []string{kv.Key}, nil
}
