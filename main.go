package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	// Prefix is the first marker of the prompt, at the beginning of the line.
	Prefix = ":: "
	// Postfix is the last marker of the prompt, just before the user input area.
	Postfix = "# "
)

// Runtime models the runtime configuration.
type Runtime struct {
	Git      bool
	Path     bool
	Verbose  bool
	ExitCode int
}

// WorkingDirectoryStatus provides the relative name of the working directory, with
// a abbreviation of the home directory to "~" if the working directory
// is the home directory.
func WorkingDirectoryStatus() string {
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

// GitStatus provides the current directory Git status.
func GitStatus(r Runtime) string {
	s := ""
	changes := len(strings.Split(execute("git status -s"), "\n"))
	if changes >= 1 {
		s = "(x)"
	}

	return s
}

// ExitStatus provides non-zero exit code status.
func ExitStatus(r Runtime) string {
	s := ""
	if r.ExitCode > 0 {
		s = "(!)"
	}

	return s
}

// BuildPrompt builds the prompt string to be printed.
func BuildPrompt(r *Runtime) string {
	prompt := ""

	// First element; right hand side padding only.
	if r.ExitCode > 0 {
		prompt = ExitStatus(*r) + " "
	}

	if r.Path {
		prompt = prompt + WorkingDirectoryStatus()
	}

	// Last element; left hand side padding only.
	if r.Git {
		prompt = prompt + " " + GitStatus(*r)
	}

	return fmt.Sprintf("%s%s%s", Prefix, prompt, Postfix)
}

func execute(cmd string) string {
	if cmd == "" {
		return ""
	}

	run := exec.Command(cmd)
	output, _ := run.CombinedOutput()
	return string(output)
}

func main() {
	g := flag.Bool("g", false, "Show Git status")
	p := flag.Bool("p", true, "Show directory path element")
	v := flag.Bool("v", false, "Verbose mode")
	x := flag.Int("x", 0, "Exit code, usually noted by the shell variable $?")
	flag.Parse()

	r := Runtime{
		Git:      *g,
		Path:     *p,
		Verbose:  *v,
		ExitCode: *x,
	}
	fmt.Print(BuildPrompt(&r))
}
