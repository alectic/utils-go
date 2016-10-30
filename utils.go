package utils

import (
	"io/ioutil"
	"os"
)

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

// CountDirs counts directories in a directory
func CountDirs(dir string) (int, error) {
	var count int
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	for _, file := range files {
		// count only directories
		if file.IsDir() {
			count++
		}
	}
	return count, nil
}
