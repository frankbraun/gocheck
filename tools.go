package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func goimports(paths []string, verbose bool) error {
	args := []string{"-l", "-w"}
	for _, path := range paths {
		args = append(args, path)
	}
	if verbose {
		fmt.Println("goimports", strings.Join(args, " "))
	}
	cmd := exec.Command("goimports", args...)
	var outbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	if outbuf.String() != "" {
		return fmt.Errorf("goimports -l -w:\n%s", strings.TrimSpace(outbuf.String()))
	}
	return nil
}

func gofmt(paths []string, verbose bool) error {
	args := []string{"-l", "-w", "-s"}
	for _, path := range paths {
		args = append(args, path)
	}
	if verbose {
		fmt.Println("gofmt", strings.Join(args, " "))
	}
	cmd := exec.Command("gofmt", args...)
	var outbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	if outbuf.String() != "" {
		return fmt.Errorf("gofmt -l -w -s:\n%s", strings.TrimSpace(outbuf.String()))
	}
	return nil
}

func golint(paths []string, verbose bool) error {
	var pathsWithSubDirs []string
	for _, path := range paths {
		fi, err := os.Stat(path)
		if err != nil {
			return err
		}
		if fi.IsDir() {
			pathsWithSubDirs = append(pathsWithSubDirs, filepath.Join(path, "..."))
		} else {
			pathsWithSubDirs = append(pathsWithSubDirs, path)
		}

	}
	if verbose {
		fmt.Println("golint", strings.Join(pathsWithSubDirs, " "))
	}
	cmd := exec.Command("golint", pathsWithSubDirs...)
	var outbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	if outbuf.String() != "" {
		return fmt.Errorf("golint:\n%s", strings.TrimSpace(outbuf.String()))
	}
	return nil
}

func govet(paths []string, verbose bool) error {
	if verbose {
		fmt.Println("go vet", strings.Join(paths, " "))
	}
	args := []string{"vet"}
	args = append(args, paths...)
	cmd := exec.Command("go", args...)
	var errbuf bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = &errbuf
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s: go vet:\n%s", err,
			strings.TrimSpace(errbuf.String()))
	}
	return nil
}

func gotest(path string, coverage, verbose bool) error {
	args := []string{"test"}
	if coverage {
		args = append(args, "-cover")
	}
	if verbose {
		args = append(args, "-v")
	}
	args = append(args, path)
	if verbose {
		fmt.Println("go", strings.Join(args, " "))
	}
	cmd := exec.Command("go", args...)
	var errbuf bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = &errbuf
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s: go test:\n%s", err,
			strings.TrimSpace(errbuf.String()))
	}
	return nil
}
