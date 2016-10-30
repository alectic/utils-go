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

// IsExistFile returns true only if file exists and is not a dir
func IsExistFile(file string) bool {
	f, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	if f.Mode().IsDir() {
		return false
	}

	return true
}

// IsExistProc returns an error if there is no process with pid
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
