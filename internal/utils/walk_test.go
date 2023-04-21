package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/ilius/is/v2"
	"github.com/sirupsen/logrus"
)

var dummyData = []byte{0, 1, 2}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	Log.Level = logrus.DebugLevel
}

type walkResult struct {
	output []string
	root   string
}

func (w *walkResult) accumulate(path string, info os.FileInfo, err error) error {
	if err == nil {
		relPath, relErr := filepath.Rel(w.root, path)
		if relErr != nil {
			return relErr
		}
		w.output = append(w.output, relPath)
	}
	return err
}

func (w *walkResult) sorted() []string {
	output := w.output
	sort.Strings(output)
	return output
}

func newWalkResult(root string) *walkResult {
	return &walkResult{root: root}
}

func mktemp(t *testing.T) string {
	is := is.New(t)
	tempDir, err := ioutil.TempDir("", "walk_test")
	is.NotErr(err)
	is.NotEqual("", tempDir)
	return tempDir
}

func TestWalkFilterBlacklist(t *testing.T) {
	is := is.New(t)
	tempDir := mktemp(t)
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()

	blackListedFilename := filepath.Join(tempDir, "ios")
	is.NotErr(ioutil.WriteFile(blackListedFilename, dummyData, 0644))

	blackListedDir := filepath.Join(tempDir, ".git")
	is.NotErr(os.Mkdir(blackListedDir, 0755))
	blackListedSubFilename := filepath.Join(blackListedDir, "config")
	is.NotErr(ioutil.WriteFile(blackListedSubFilename, dummyData, 0644))

	whiteListedFilename := filepath.Join(tempDir, "whiteListedFile.dat")
	is.NotErr(ioutil.WriteFile(whiteListedFilename, dummyData, 0644))

	whiteListedDirectory := filepath.Join(tempDir, "whiteListedDir")
	is.NotErr(os.Mkdir(whiteListedDirectory, 0755))
	whiteListedSubFilename := filepath.Join(whiteListedDirectory, "whiteListedSubFilename")
	is.NotErr(ioutil.WriteFile(whiteListedSubFilename, dummyData, 0644))

	result := newWalkResult(tempDir)
	is.NotErr(filepath.Walk(tempDir, WalkFilterBlacklist(tempDir, result.accumulate)))
	output := result.sorted()
	is.Equal(len(output), 4)
	is.Equal(".", output[0])
	is.Equal("whiteListedDir", output[1])
	is.Equal(filepath.Join("whiteListedDir", "whiteListedSubFilename"), output[2])
	is.Equal("whiteListedFile.dat", output[3])
}

func createSimpleFilesystem(tempDir string, t *testing.T) {
	is := is.New(t)
	file := filepath.Join(tempDir, "file.txt")
	is.NotErr(ioutil.WriteFile(file, dummyData, 0644))

	dir := filepath.Join(tempDir, "dir.dirext")
	is.NotErr(os.Mkdir(dir, 0755))
	subFile := filepath.Join(dir, "subfile.png")
	is.NotErr(ioutil.WriteFile(subFile, dummyData, 0644))
}

func TestWalkOnlyDirectory(t *testing.T) {
	is := is.New(t)
	tempDir := mktemp(t)
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()
	createSimpleFilesystem(tempDir, t)

	result := newWalkResult(tempDir)
	is.NotErr(filepath.Walk(tempDir, WalkOnlyDirectory(result.accumulate)))
	output := result.sorted()
	is.Equal(len(output), 2)
	is.Equal(".", output[0])
	is.Equal("dir.dirext", output[1])
}

func TestWalkOnlyFile(t *testing.T) {
	is := is.New(t)
	tempDir := mktemp(t)
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()
	createSimpleFilesystem(tempDir, t)

	result := newWalkResult(tempDir)
	is.NotErr(filepath.Walk(tempDir, WalkOnlyFile(result.accumulate)))
	output := result.sorted()
	is.Equal(len(output), 2)
	is.Equal(filepath.Join("dir.dirext", "subfile.png"), output[0])
	is.Equal("file.txt", output[1])
}

func TestWalkFilterPrefix(t *testing.T) {
	is := is.New(t)
	tempDir := mktemp(t)
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()
	createSimpleFilesystem(tempDir, t)

	result := newWalkResult(tempDir)
	is.NotErr(filepath.Walk(tempDir, WalkFilterPrefix(result.accumulate, "dir")))
	output := result.sorted()
	is.Equal(len(output), 2)
	is.Equal(".", output[0])
	is.Equal("file.txt", output[1])
}

func TestWalkOnlyExtension(t *testing.T) {
	is := is.New(t)
	tempDir := mktemp(t)
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()
	createSimpleFilesystem(tempDir, t)

	result := newWalkResult(tempDir)
	is.NotErr(filepath.Walk(tempDir, WalkOnlyExtension(result.accumulate, "txt")))
	output := result.sorted()
	is.Equal(len(output), 3)
	is.Equal(".", output[0])
	is.Equal("dir.dirext", output[1])
	is.Equal("file.txt", output[2])
}
