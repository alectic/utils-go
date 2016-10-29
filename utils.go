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
