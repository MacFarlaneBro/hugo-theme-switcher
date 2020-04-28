package hugothemeswitcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git"
)

const themes string = "https://github.com/gohugoio/hugoThemes"
const themesDirectory string = "./themes"

func main() {
	getHugoTheme()
	fmt.Println("Welcome to the hugo theme switcherdafd.")
}

func getHugoTheme() {
	//pull latest version of hugo themes repo
	if _, err := os.Stat(themesDirectory); os.IsNotExist(err) {

		_, err := git.PlainClone(themesDirectory, false, &git.CloneOptions{
			URL: themes,
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

	// Remove all non-url lines
	urlFilter := func(s string) bool { return strings.HasPrefix(s, "url") }

	// add all themes to a slice
	urls := filter(lines, urlFilter)

	// select a random repo
	rand.Seed(time.Now().Unix())
	// print it out
	fmt.Println(urls[rand.Intn(len(urls))])
	// download the theme
}

//Filter function lifted from here: https://stackoverflow.com/questions/37562873/most-idiomatic-way-to-select-elements-from-an-array-in-golang#answer-37563128
// Given an array of strings return an array of only those strings which are true according to the supplied 'test' function
func filter(lines []string, test func(string) bool) (ret []string) {
	for _, s := range lines {
		url := strings.TrimSpace(s)
		if test(url) {
			ret = append(ret, url[6:])
		}
	}
	return
}
