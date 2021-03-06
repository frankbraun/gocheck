package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var toolPaths = []string{
	"golang.org/x/lint/golint",
	"golang.org/x/tools/cmd/goimports",
}

func getTools(verbose bool) error {
	for _, tool := range toolPaths {
		if err := getTool(tool, verbose); err != nil {
			return err
		}
	}
	return nil
}

func getTool(tool string, verbose bool) error {
	args := []string{"get"}
	if verbose {
		args = append(args, "-v")
	}
	args = append(args, tool)
	if verbose {
		fmt.Println("go", strings.Join(args, " "))
	}
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
