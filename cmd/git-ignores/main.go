package main

import (
	flag "github.com/spf13/pflag"
	"os"
	"io"
	"fmt"
	"net/http"
)

var (
	templateNameFlag string
	forceFlag bool
	appendFlag bool
)

func init() {
	flag.StringVarP(&templateNameFlag, "template", "t", "", "Gitignore template.")
	flag.BoolVarP(&forceFlag, "force", "f", false, "Force overwrite existing .gitignore.")
	flag.BoolVarP(&appendFlag, "append", "a", false, "Tack Gitignore template onto the end of existing .gitignore file.")
}

func main() {
	// Parse cli flag options
	flag.Parse()

	// URL for .gitignore file
	if templateNameFlag == "" {
		fmt.Printf("ERROR: Please provide a template via --template\n")
		os.Exit(1)
	}
	templateURL := fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/main/%s.gitignore", templateNameFlag)

	// Make GET request to templateURL
	res, err := http.Get(templateURL)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(bodyBytes))
	}
	
}
