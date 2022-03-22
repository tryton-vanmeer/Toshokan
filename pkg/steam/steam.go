package steam

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Jleagle/steam-go/steamvdf"
)

const (
	STEAM_APPS_ROOT = ".steam/steam/steamapps"
)

var FILTER_LIST = []string{
	"Proton",
	"Steam Linux Runtime",
	"Steamworks Common Redistributables",
}

type App struct {
	Name             string
	AppID            string
	LibraryFolder    string
	InstallDirectory string
}

type AppList []App

func (apps AppList) Len() int {
	return len(apps)
}

func (apps AppList) Less(i, j int) bool {
	return apps[i].Name < apps[j].Name
}

func (apps AppList) Swap(i, j int) {
	apps[i], apps[j] = apps[j], apps[i]
}

func (apps AppList) Sort() {
	sort.Sort(apps)
}

// get the URL for the apps page on Steam
func (app App) GetStorePage() string {
	return fmt.Sprintf("https://store.steampowered.com/app/%s", app.AppID)
}

// get path game would use for it's proton prefix
func (app App) ProtonPrefix() string {
	return fmt.Sprintf("%s/compatdata/%s", app.LibraryFolder, app.AppID)
}

// check if game uses proton
func (app App) IsProton() bool {
	// if the proton prefix path exists, assume it is using proton
	// game can still be using proton, but hasn't performed first launch (so path won't exist)
	_, err := os.Stat(app.ProtonPrefix())

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

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	installDirectory := strings.Join([]string{libraryFolder, "common", app_manifest["installdir"]}, "/")
	installDirectory = strings.Replace(installDirectory, home, "~", 1)

	return App{
		Name:             app_manifest["name"],
		AppID:            app_manifest["appid"],
		LibraryFolder:    libraryFolder,
		InstallDirectory: installDirectory,
	}
}

// filter out tools and runtimes
func (apps AppList) filter() (new AppList) {
	for _, app := range apps {
		flagged := false
		for _, filter := range FILTER_LIST {
			if strings.Contains(app.Name, filter) {
				flagged = true
				break
			}
		}

		if !flagged {
			new = append(new, app)
		}
	}

	return
}

// return a list of the installed apps
func GetApps() (apps AppList) {
	folders := libraryFolders()

	// search library directories for app manifest files
	for _, folder := range folders {
		files, err := ioutil.ReadDir(folder)

		if err != nil {
			log.Fatalf("error reading directory: %s", folder)
		}

		for _, file := range files {
			if strings.Contains(file.Name(), "appmanifest") {
				app := parseAppManifest(folder, file.Name())
				apps = append(apps, app)
			}
		}
	}

	return apps.filter()
}
