package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	// MaxWordLen specifies the maximum word length in sub-parts.
	MaxWordLen = 20
	// ShellMarker provides a visual clue of a new shell line.
	ShellMarker = "::"
	// GitDiffMarker signals Git diff in the current repository.
	GitDiffMarker = "!"
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

func getMarker(exitCode int) string {
	// Highly opinionated shell marker: two colons in green if exitCode is zero,
	// otherwise to colons in red.
	if exitCode == 0 {
		return fmt.Sprintf("\033[0;32m%s\033[0m", ShellMarker)
	}
	return fmt.Sprintf("\033[0;7m%s\033[0m", ShellMarker)
}

func getPwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		return ""
	}

	if os.Getenv("TELEPROMPT_DISABLE_WORD_LEN") == "" {
		if utf8.RuneCountInString(pwd) > MaxWordLen {
			pwd = filepath.Base(pwd)
		}

		if utf8.RuneCountInString(pwd) > MaxWordLen {
			return fmt.Sprintf(
				"%s",
				shorten(pwd, 10, true))
		}
	}
	return pwd
}

func getGitBranchName() string {
	var (
		output []byte
		err    error
		cmd    string
		name   string
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
	// Add status; postfix "!" if files changed or added.
	name = strings.TrimSpace(string(output))

	statusArgs := []string{"ls-files", "-mo"}
	if output, err = exec.Command(cmd, statusArgs...).Output(); err == nil {
		name = name + GitDiffMarker
	}
	return name

}

// BuildPrompt assembles the prompt contents.
func BuildPrompt(exitCode int) string {
	return fmt.Sprintf("%s %s %s# ", getMarker(exitCode), getGitBranchName(), getPwd())
}

func main() {
	exitCode := 0
	if len(os.Args) > 1 {
		if i, err := strconv.Atoi(os.Args[1]); err == nil {
			exitCode = i
		}
	}
	fmt.Printf("%s", BuildPrompt(exitCode))
}
