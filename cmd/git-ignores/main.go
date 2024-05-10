package main

import (
	flag "github.com/spf13/pflag"
	"os"
	"io"
	"errors"
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

		// Create string from GET result
		var gitignoreContents string = string(bodyBytes)

		// Check if .gitignore exists
		if _, err := os.Stat("./.gitignore"); err != nil {
			if forceFlag {
				os.Remove(".gitignore")
				f, err := os.Create(".gitignore")
				if err != nil {
					panic(err)
				}
				defer f.Close()
				_,err = f.WriteString(gitignoreContents)
				if err != nil {
					panic(err)
				}
			}
		} else if errors.Is(err, os.ErrNotExist) {
		    f, err := os.Create(".gitignore")
		    if err != nil {
			    panic(err)
		    }
		    defer f.Close()
		    _,err = f.WriteString(gitignoreContents)
		    if err != nil {
			    panic(err)
		    }

		}
	}
	
}
