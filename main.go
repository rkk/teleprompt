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
	// GitMarker marks a dirty Git state.
	GitMarker = "(x)"
	// ExitMarker marks a non-zero exit code.
	ExitMarker = "(!)"
	// ApplicationVersion is the version of this application
	ApplicationVersion = "0.0.2"
)

// Runtime models the runtime configuration.
type Runtime struct {
	Git      bool
	Path     bool
	Verbose  bool
	ExitCode int
}

// WorkingDirectoryStatus provides the relative name of the working directory.
func WorkingDirectoryStatus() string {
	var dir string
	var parts []string
	home := os.Getenv("HOME")
	cwd, err := os.Getwd()

	if err != nil || cwd == home {
		return ""
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
		s = GitMarker
	}

	return s
}

// ExitStatus provides non-zero exit code status.
func ExitStatus(r Runtime) string {
	s := ""
	if r.ExitCode > 0 {
		s = ExitMarker
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

func showHelp() string {
	return `usage: teleprompt OPTIONS

OPTIONS:
-g     Include Git status
-h     Display this help text and exit
-p     Include directory path element
-v     Display application version and exit
-V     Run in verbose mode, displaying more information
-x ARG Include exit code, where ARG is usually noted by the shell variable $?
`
}

func main() {
	g := flag.Bool("g", false, "Show Git status")
	help := flag.Bool("h", false, "Show help instructions")
	p := flag.Bool("p", false, "Show directory path element")
	verbose := flag.Bool("V", false, "Verbose mode")
	version := flag.Bool("v", false, "Show application version")
	x := flag.Int("x", 0, "Exit code, usually noted by the shell variable $?")
	flag.Parse()

	if *version {
		fmt.Println(ApplicationVersion)
		os.Exit(0)
	}

	if *help || (!*g && !*p && *x == 0) {
		fmt.Println(showHelp())
		os.Exit(0)
	}

	r := Runtime{
		Git:      *g,
		Path:     *p,
		Verbose:  *verbose,
		ExitCode: *x,
	}
	fmt.Print(BuildPrompt(&r))
}
