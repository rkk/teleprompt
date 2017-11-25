package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

const (
	// MaxWordLen specifies the maximum word length in sub-parts.
	MaxWordLen = 20
)

func shorten(s string, count int, fromStart bool) string {
	runes := []rune(s)
	if fromStart {
		return fmt.Sprintf(
			"%s",
			string(runes[0:count]))
	}
	return fmt.Sprintf(
		"%s",
		string(runes[len(runes)-count:len(runes)]))
}

func getMarker() string {
	// Highly opinionated shell marker: two colons in green.
	return "\033[0;32m::\033[0m"
}

func getPwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		return ""
	}

	if utf8.RuneCountInString(pwd) > MaxWordLen {
		pwd = filepath.Base(pwd)
	}

	if utf8.RuneCountInString(pwd) > MaxWordLen {
		return fmt.Sprintf(
			"%s",
			shorten(pwd, 10, true))
	}
	return pwd
}

func getGitBranchName() string {
	var (
		output []byte
		err    error
		cmd    string
	)

	gitDir, err := os.Getwd()

	if err != nil {
		return ""
	}
	if _, err := os.Stat(gitDir + "/.git"); os.IsNotExist(err) {
		return ""
	}

	cmd = "git"
	args := []string{"rev-parse", "--abbrev-ref", "HEAD"}
	if output, err = exec.Command(cmd, args...).Output(); err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))

}

// BuildPrompt assembles the prompt contents.
func BuildPrompt() string {
	return fmt.Sprintf("%s %s %s# ", getMarker(), getGitBranchName(), getPwd())
}

func main() {
	fmt.Printf("%s", BuildPrompt())
}
