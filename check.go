package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func check(paths, exclude []string, coverage, verbose bool) error {
	excludeMap := make(map[string]struct{})
	for _, ex := range exclude {
		excludeMap[ex] = struct{}{}
	}
	var (
		files        []string
		testDirs     []string
		testFileDirs []string
	)
	for _, path := range paths {
		switch fi, err := os.Stat(path); {
		case err != nil:
			return err
		case fi.IsDir():
			if err := checkDir(path, excludeMap, verbose); err != nil {
				return err
			}
			td, err := getTestDirs(path, excludeMap)
			if err != nil {
				return err
			}
			testDirs = append(testDirs, td...)
		default:
			files = append(files, path)
			if strings.HasSuffix(path, "_test.go") {
				testFileDirs = append(testFileDirs, filepath.Base(path))
			}
		}
	}
	if len(files) > 0 {
		if err := checkFiles(files, verbose); err != nil {
			return err
		}
	}
	// go test
	fileDirs := make(map[string]struct{})
	for _, d := range testFileDirs {
		fileDirs[d] = struct{}{}
	}
	var dirs []string
	for d := range fileDirs {
		dirs = append(dirs, d)
	}
	sort.Strings(dirs)
	for _, d := range dirs {
		if err := gotest(d, coverage, verbose); err != nil {
			return err
		}
	}
	for _, d := range testDirs {
		err := gotest("."+string(filepath.Separator)+d, coverage, verbose)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkDir(path string, excludeMap map[string]struct{}, verbose bool) error {
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	var files []string
	for _, fi := range fis {
		fn := fi.Name()
		// ignore hidden directories and files
		if !strings.HasPrefix(fn, ".") {
			if fi.IsDir() {
				_, ok := excludeMap[fn]
				if !ok {
					err := checkSubdir(filepath.Join(path, fn), verbose)
					if err != nil {
						return err
					}
				}
			} else if strings.HasSuffix(fn, ".go") {
				files = append(files, fn)
			}
		}
	}
	if len(files) > 0 {
		if err := checkFiles(files, verbose); err != nil {
			return err
		}
	}
	return nil
}

func checkFiles(files []string, verbose bool) error {
	if err := goimports(files, verbose); err != nil {
		return err
	}
	if err := gofmt(files, verbose); err != nil {
		return err
	}
	if err := golint(files, verbose); err != nil {
		return err
	}
	if err := gotoolvet(files, verbose); err != nil {
		return err
	}
	return nil
}

func checkSubdir(subdir string, verbose bool) error {
	path := []string{subdir}
	if err := goimports(path, verbose); err != nil {
		return err
	}
	if err := gofmt(path, verbose); err != nil {
		return err
	}
	if err := golint(path, verbose); err != nil {
		return err
	}
	if err := gotoolvet(path, verbose); err != nil {
		return err
	}
	return nil
}

func getTestDirs(path string, excludeMap map[string]struct{}) ([]string, error) {
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var (
		dirs    []string
		subDirs []string
		include bool
	)
	for _, fi := range fis {
		fn := fi.Name()
		// ignore hidden directories and files
		if !strings.HasPrefix(fn, ".") {
			if fi.IsDir() {
				_, ok := excludeMap[fn]
				if !ok {
					d, err := getTestDirs(filepath.Join(path, fn), excludeMap)
					if err != nil {
						return nil, err
					}
					subDirs = append(subDirs, d...)
				}
			} else if !include && strings.HasSuffix(fn, "_test.go") {
				include = true
			}
		}
	}
	if include {
		dirs = append(dirs, path)
	}
	dirs = append(dirs, subDirs...)
	return dirs, nil
}
