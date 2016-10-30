package utils

import (
	"io/ioutil"
	"os"
)

type FSCount struct {
	Files int
	Dirs  int
	All   int
}

// IsExistFile returns an error only if the file doesn't exist
func IsExistFile(file string) error {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return err
		}
	}
	return nil
}

// IsExistProc returns an error if there is no process with specified pid
func IsExistProcByPid(pid int) error {
	_, err := os.FindProcess(pid)
	return err
}

// CountInDir counts the number of items in dir and returns a FSCount struct
// with the number of dirs, files and total number of items
func CountInDir(dir string) (FSCount, error) {
	count := FSCount{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return count, err
	}

	for _, file := range files {
		count.All += 1
		if file.Mode().IsRegular() {
			count.Files += 1
		} else if file.Mode().IsDir() {
			count.Dirs += 1
		}
	}

	return count, nil
}
