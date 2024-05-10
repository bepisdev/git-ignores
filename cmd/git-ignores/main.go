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

func fetchGitignoreTemplate(templateName string) (string,error) {
	templateURL := fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/main/%s.gitignore", templateNameFlag)
	res,err := http.Get(templateURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code %d", res.StatusCode)
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}

func main() {
	// Parse cli flag options
	flag.Parse()

	// URL for .gitignore file
	if templateNameFlag == "" {
		fmt.Fprintf(os.Stderr, "ERROR: Please provide a template via --template\n")
		flag.Usage()
		os.Exit(1)
	}

	gitignoreContents, err := fetchGitignoreTemplate(templateNameFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to fetch template: %v\n", templateNameFlag)
		os.Exit(1)
	}

	// Check if .gitignore exists
	if _,err := os.Stat(".gitignore"); err == nil {
		if forceFlag {
			if err := os.Remove(".gitignore"); err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: Failed to remove existing .gitignore file\n")
				os.Exit(1)
			}
		} else {
			fmt.Fprintf(os.Stderr, "ERROR: .gitignore already exists. Use --force to overwrite\n")
			flag.Usage()
			os.Exit(1)
		}
	} 

	// Write contents to .gitignore
	if err := writeIgnoreFile(gitignoreContents); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to write .gitignore: %v\n", err)
			os.Exit(1)
	}
}
