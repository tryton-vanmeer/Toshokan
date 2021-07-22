package steam

import "fmt"

var STEAM_LIBRARY_FOLDERS_VDF = ".steam/steam/steamapps/libraryfolders.vdf"

func getLibraryFolders() []string {
	return []string{"hello"}
}

func Test() {
	folders := getLibraryFolders()

	fmt.Println(folders)
}
