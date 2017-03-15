package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

var (
	// ErrSliceMismatch is returned when slices are of different lengths
	ErrSliceMismatch = errors.New("slices are not the same length")
	// ErrSlicesEmpty is returned when slices are empty
	ErrSlicesEmpty = errors.New("slices are empty")
	// ErrSliceEmpty is returned when a slice is empty
	ErrSliceEmpty = errors.New("slice is empty")
)

// FSCount holds the count of Files Directories and Their Total
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

// IsExistDir return true only if dir exists and is not a file
func IsExistDir(dir string) bool {
	f, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	if f.Mode().IsRegular() {
		return false
	}

	return true
}

// IsExistProcPid returns false if there is no process with pid
// otherwise returns true if the process exist
func IsExistProcPid(pid int) bool {
	p, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	if err := p.Signal(syscall.Signal(0)); err != nil {
		return false
	}

	return true
}

// IsExistProcName return false if there is no process with name
// otherwise returns true if the process exists
func IsExistProcName(name string) bool {
	procPath := "/proc"
	var ok bool
	entries, _ := ioutil.ReadDir(procPath)

	for _, entry := range entries {
		if !entry.Mode().IsDir() {
			continue
		}

		// if dir's name is not made of digits
		if _, err := strconv.Atoi(entry.Name()); err != nil {
			continue
		}

		b, _ := ioutil.ReadFile(filepath.Join(procPath, entry.Name(), "comm"))

		if strings.Contains(string(b), name) {
			ok = true // process exists
			break
		}
	}

	return ok
}

// CountDir counts the number of items in dir and returns a FSCount struct
// with the number of dirs, files and total number of entries
func CountDir(dir string) (FSCount, error) {
	count := FSCount{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return count, err
	}

	for _, file := range files {
		count.All++
		if file.Mode().IsRegular() {
			count.Files++
		} else if file.Mode().IsDir() {
			count.Dirs++
		}
	}

	return count, nil
}

// Zip takes two slices of the same length, compares their length
// and returns an error if one of them or both are empty
// or if they mismatch otherwise returns a map
func Zip(s1, s2 []string) (map[string]string, error) {
	if len(s1) == 0 && len(s2) == 0 {
		return nil, ErrSlicesEmpty
	}

	if len(s1) == 0 || len(s2) == 0 {
		return nil, ErrSliceEmpty
	}

	if len(s1) != len(s2) {
		return nil, ErrSliceMismatch
	}

	m := make(map[string]string)

	for i := 0; i < len(s1); i++ {
		m[s1[i]] = s2[i]
	}

	return m, nil
}
