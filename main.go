package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	prefix  = ":: "
	postfix = "# "
)

// WorkingDirectory provides the relative name of the working directory, with
// a abbreviation of the home directory to "~" if the working directory
// is the home directory.
func WorkingDirectory() string {
	var dir string
	var parts []string
	home := os.Getenv("HOME")
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}

	if cwd == home {
		return "~"
	}

	if cwd == "/" {
		return cwd
	}

	parts = strings.Split(cwd, string(os.PathSeparator))
	if len(parts) > 1 {
		dir = parts[len(parts)-1]
	}

	return dir
}

// BuildPrompt builds the prompt string to be printed.
func BuildPrompt() string {
	return fmt.Sprintf("%s%s%s", prefix, WorkingDirectory(), postfix)
}

func main() {
	fmt.Print(BuildPrompt())
}
