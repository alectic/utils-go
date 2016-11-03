package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
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

// CountInDir counts the number of items in dir and returns a FSCount struct
// with the number of dirs, files and total number of entries
func CountDir(dir string) (FSCount, error) {
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
