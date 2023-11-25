package builders

import (
	"fmt"
	"os/exec"
	"strings"
)

var CurrentDirectory string

func CreateDirectories() {
	pages, _, _ := getDirectoriesPaths()

	fmt.Println(pages)

	// errPages := os.Mkdir(pages, 0644)
	// if errPages != nil {
	// 	panic(errPages)
	// }
	// errModules := os.Mkdir(modules, 0644)
	// if errModules != nil {
	// 	panic("Failed to create Modules directory")
	// }

	// errFunctions := os.Mkdir(functions, 0644)
	// if errFunctions != nil {
	// 	panic("Failed to create Function directory")
	// }

}

func getDirectoriesPaths() (string, string, string) {
	pwd := exec.Command("pwd")
	output, err := pwd.Output()
	if err != nil {
		panic("Failed to retrieve current directory")
	}

	CurrentDirectory = string(output)
	pages := strings.Split(CurrentDirectory, `
	`)

	pagesDir := fmt.Sprintf("%v", pages[0])
	pagesModules := fmt.Sprintf("%s/modules", CurrentDirectory)
	pagesFunctions := fmt.Sprintf("%s/functions", CurrentDirectory)

	return pagesDir, pagesModules, pagesFunctions
}
