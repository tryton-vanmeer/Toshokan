package steam

import "fmt"

var STEAM_APPS_ROOT = ".steam/steam/steamapps"
var STEAM_LIBRARY_FOLDERS_VDF = fmt.Sprintf("%s/libraryfolders.vdf", STEAM_APPS_ROOT)

func getLibraryFolders() []string {
	return []string{"hello"}
}

func Test() {
	folders := getLibraryFolders()

	fmt.Println(folders)
}
