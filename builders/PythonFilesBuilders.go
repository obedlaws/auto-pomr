package builders

import (
	"os"
	"pomr/files"
)

func CreateFiles() {
	err := os.WriteFile("pages/home_page.py", files.HomePageObjects, 0644)
	if err != nil {
		panic(err)
	}

}
