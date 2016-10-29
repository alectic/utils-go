package utils

import "os"

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
