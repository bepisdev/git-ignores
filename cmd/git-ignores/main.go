package main

import (
	flag "github.com/spf13/pflag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	templateName string
	outputPath string
	forceFlag    bool
)

func init() {
	flag.StringVarP(&templateName, "template", "t", "", "Gitignore template.")
	flag.StringVarP(&outputPath, "output", "o", ".gitignore", "Path to write .gitignore contents to.")
	flag.BoolVarP(&forceFlag, "force", "f", false, "Force overwrite existing .gitignore.")
}

func writeIgnoreFile(content string) error {
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	return err
}

func fetchGitignoreTemplate(templateName string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	templateURL := fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/main/%s.gitignore", templateName)
	res, err := client.Get(templateURL)
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
	flag.Parse()

	if templateName == "" {
		fmt.Fprintf(os.Stderr, "ERROR: Please provide a template via --template\n")
		flag.Usage()
		os.Exit(1)
	}

	gitignoreContents, err := fetchGitignoreTemplate(templateName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to fetch template: %v\n", err)
		os.Exit(1)
	}

	if _, err := os.Stat(".gitignore"); err == nil && !forceFlag {
		fmt.Fprintf(os.Stderr, "ERROR: .gitignore already exists. Use --force to overwrite\n")
		flag.Usage()
		os.Exit(1)
	}

	if forceFlag {
		if err := os.Remove(".gitignore"); err != nil && !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to remove existing .gitignore file: %v\n", err)
			os.Exit(1)
		}
	}

	if err := writeIgnoreFile(gitignoreContents); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to write .gitignore: %v\n", err)
		os.Exit(1)
	}
}
