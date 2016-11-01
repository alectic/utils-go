package utils_test

import (
	"math/rand"
	"os"
	"path"
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
	assert.NoError(os.Mkdir(testdir, 0744), "should not return an error")
	defer func() {
		assert.NoError(os.RemoveAll(testdir), "should not return an error")
	}()
	assert.True(utils.IsExistDir(testdir), "should return true")
}

func TestIsExistProcByPid(t *testing.T) {
	pid := os.Getpid()
	assert.NoError(t, utils.IsExistProcByPid(pid), "should not retun an error")
}

func TestCountInDir(t *testing.T) {
	filename := "file_test"
	dirname := "dir_test"
	testdir := path.Join(os.TempDir(), "utils-go.test")

	assert := assert.New(t)
	require := require.New(t)

	require.NoError(os.RemoveAll(testdir), "should not return an error")
	require.NoError(os.Mkdir(testdir, 0744), "should not return an error")

	maxItems := 20
	var files, dirs int
	choices := map[int]func(int){
		0: func(i int) {
			_, err := os.Create(path.Join(testdir, filename+strconv.Itoa(i)))
			require.NoError(err, "should not return an error")
			files += 1
		},
		1: func(i int) {
			err := os.Mkdir(path.Join(testdir, dirname+strconv.Itoa(i)), 0744)
			require.NoError(err, "should not return an error")
			dirs += 1
		},
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i < maxItems+1; i++ {
		choices[rng.Intn(2)](i)
	}

	defer func() {
		assert.NoError(os.RemoveAll(testdir), "shout not return an error")
	}()

	count, err := utils.CountInDir(testdir)
	require.NoError(err, "should not return an error")
	assert.Equal(files, count.Files, "they should be equal")
	assert.Equal(dirs, count.Dirs, "they should be equal")
	assert.Equal(maxItems, count.All, "they should be equal")
}
