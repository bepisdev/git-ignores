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
	gitignoreContents string
)

func init() {
	flag.StringVarP(&templateNameFlag, "template", "t", "", "Gitignore template.")
	flag.BoolVarP(&forceFlag, "force", "f", false, "Force overwrite existing .gitignore.")
}

func writeIgnoreFile(cont string) error {
	f, err := os.Create(".gitignore")
	if err != nil {
		return err
	}
	defer f.Close()
	_,err = f.WriteString(cont)
	if err != nil {
		return err
	}
	return nil

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
		gitignoreContents = string(bodyBytes)

	} else {
		fmt.Println("HTTP Request failure")
		os.Exit(1)
	}

	// Check if .gitignore exists
	if _,err := os.Stat(".gitignore"); err == nil {
		if forceFlag {
			os.Remove(".gitignore")
		} else {
			fmt.Println(".gitignore already exists, use --force to overwrite")
			os.Exit(1)
		}
	} 
	err = writeIgnoreFile(gitignoreContents)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
