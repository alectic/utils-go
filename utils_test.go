package utils_test

import (
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	utils "github.com/alectic/utils-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsExistFile(t *testing.T) {
	testfile := "/tmp/utils-go.testfile"
	assert := assert.New(t)
	assert.NoError(os.RemoveAll(testfile), "should not return an error")
	assert.False(utils.IsExistFile(testfile), "should return false")
	_, err := os.Create(testfile)
	assert.NoError(err, "should not return an error")
	defer func() {
		assert.NoError(os.Remove(testfile), "should not return an error")
	}()
	assert.True(utils.IsExistFile(testfile), "should return true")
}

func TestIsExistDir(t *testing.T) {
	testdir := "/tmp/utils-go.testdir"
	assert := assert.New(t)
	assert.NoError(os.RemoveAll(testdir), "shoult not return an error")
	assert.False(utils.IsExistDir(testdir), "should return false")
	require.NoError(t, os.Mkdir(testdir, 0744), "should not return an error")
	defer func() {
		assert.NoError(os.RemoveAll(testdir), "should not return an error")
	}()
	assert.True(utils.IsExistDir(testdir), "should return true")
}

func TestIsExistProcPid(t *testing.T) {
	pid := os.Getpid()
	assert := assert.New(t)
	assert.True(utils.IsExistProcPid(pid), "should return true")
	assert.False(utils.IsExistProcPid(12345), "should return false")
}

func TestIsExistProcName(t *testing.T) {
	name := filepath.Base(os.Args[0])
	assert := assert.New(t)
	assert.True(utils.IsExistProcName(name), "should return true")
	assert.False(utils.IsExistProcName("name1234567890"), "should return false")
}

func TestCountDir(t *testing.T) {
	filename := "file_test"
	dirname := "dir_test"
	testdir := filepath.Join(os.TempDir(), "utils-go.test")

	assert := assert.New(t)
	require := require.New(t)

	assert.NoError(os.RemoveAll(testdir), "should not return an error")
	require.NoError(os.Mkdir(testdir, 0744), "should not return an error")

	maxEntries := 20
	var files, dirs int
	choices := map[int]func(int){
		0: func(i int) {
			_, err := os.Create(filepath.Join(testdir, filename+strconv.Itoa(i)))
			require.NoError(err, "should not return an error")
			files += 1
		},
		1: func(i int) {
			err := os.Mkdir(filepath.Join(testdir, dirname+strconv.Itoa(i)), 0744)
			require.NoError(err, "should not return an error")
			dirs += 1
		},
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i < maxEntries+1; i++ {
		choices[rng.Intn(2)](i)
	}

	defer func() {
		assert.NoError(os.RemoveAll(testdir), "shout not return an error")
	}()

	count, err := utils.CountDir(testdir)
	require.NoError(err, "should not return an error")
	assert.Equal(files, count.Files, "they should be equal")
	assert.Equal(dirs, count.Dirs, "they should be equal")
	assert.Equal(maxEntries, count.All, "they should be equal")
}
