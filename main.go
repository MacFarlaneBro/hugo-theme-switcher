package main

import (
	"strings"
	"log"
    "fmt"
    "os"
    "github.com/go-git/go-git"
    "io/ioutil"
)	

const themes string = "https://github.com/gohugoio/hugoThemes"
const themesDirectory string = "./themes"

func main() {
	getHugoTheme()
    fmt.Println("Welcome to the hugo theme switcher.")
}


func getHugoTheme() {
	//pull latest version of hugo themes repo
	if _, err := os.Stat(themesDirectory); os.IsNotExist(err) {

		_, err := git.PlainClone(themesDirectory, false, &git.CloneOptions{
			URL:      themes,
			Progress: os.Stdout,
		})

		if err != nil {
			log.Fatal(err)
		}
	}

	// The themes in this repository are actually stored as submodules on the main repo,
	// therefore the only file we actually need in order to access them is .gitmodules as
	// that contains a list of all available themes 
    modulesFile, err := ioutil.ReadFile(themesDirectory + "/.gitmodules")

	if err != nil {
		log.Fatal(err)
	}

	// Split the file out into its individual lines
	lines := strings.Split(string(modulesFile), "\n")
	fmt.Println(lines[0])
	//add all themes to channel

	// select a random repo

	// download the theme
}
