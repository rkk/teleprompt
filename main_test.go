package main

import (
	"os"
	"testing"
)

func TestGetGitHasDiffPositive(t *testing.T) {
	pwd, _ := os.Getwd()
	runtime := Runtime{GitCommand: "testhelpers/fakegit-onechange.sh", WorkingDirectory: pwd}

	diff := getGitHasDiff(runtime)
	if !diff {
		t.Errorf("Expected true, got false")
	}
}

func TestGetGitHasDiffNegative(t *testing.T) {
	pwd, _ := os.Getwd()
	runtime := Runtime{GitCommand: "testhelpers/fakegit-nochange.sh", WorkingDirectory: pwd}

	diff := getGitHasDiff(runtime)
	if diff {
		t.Errorf("Expected false, got true")
	}
}
