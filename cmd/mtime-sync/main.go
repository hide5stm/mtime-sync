package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// findMaxSizeFile  is a function that finds the maximum size file.
func findMaxSizeFile(directoryPath string) (string, time.Time, error) {
	var maxFileSize int64
	var maxFileModTime time.Time
	var maxFilePath string

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if info.Size() > maxFileSize {
				maxFileSize = info.Size()
				maxFileModTime = info.ModTime()
				maxFilePath = path
			}
		}
		return nil
	})

	if err != nil {
		return "", time.Time{}, err
	}

	if maxFilePath == "" {
		return "", time.Time{}, fmt.Errorf("no valid file found in %s", directoryPath)
	}

	return maxFilePath, maxFileModTime, nil
}

// setMtimeToDirectory is a function that sets the mtime of a directory.
func setMtimeToDirectory(directoryPaths []string, verbose bool) {
	for _, directoryPath := range directoryPaths {
		filePath, maxFileModTime, err := findMaxSizeFile(directoryPath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		err = os.Chtimes(directoryPath, maxFileModTime, maxFileModTime)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			if verbose {
				fmt.Printf("set mtime of directory '%s' to mtime of '%s' .\n", directoryPath, filePath)
			}
		}
	}
}

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "Verbose mode (output messages)")
	flag.Parse()
	directoryPaths := flag.Args()

	if len(directoryPaths) == 0 {
		fmt.Println("Usage: go run main.go [-v] directory_path1 directory_path2 ...")
		return
	}

	setMtimeToDirectory(directoryPaths, verbose)
}
