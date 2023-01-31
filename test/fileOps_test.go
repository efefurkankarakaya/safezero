package test

import (
	"os"
	"path/filepath"
	fileOperations "safezero/internal/file"
	"testing"
)

// $ go test -v ./test

func TestFileDeletion(t *testing.T) {
	const want bool = false
	const testDir string = "temp/"

	var got bool
	var message string

	dir, _ := os.Getwd()
	fileOperations.CreateTestingFolder(dir)
	fileOperations.CopyTestingFilesToTestingFolder(dir)

	filepath.Walk(testDir, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			fileOperations.RemoveFile(path)
			got = fileOperations.CheckIfPathExists(path)
		}

		return err
	})

	fileOperations.RemoveTree(testDir)

	if got != want {
		message = "File deletion test failed."
		t.Errorf(message)
	}
}

func TestTreeDeletion(t *testing.T) {
	const want bool = false
	const testDir string = "temp/"

	var got bool
	var message string

	dir, _ := os.Getwd()
	fileOperations.CreateTestingFolder(dir)
	fileOperations.CopyTestingFilesToTestingFolder(dir)

	fileOperations.RemoveTree(testDir)

	got = fileOperations.CheckIfPathExists(testDir)

	if got != want {
		message = "Could not removed tree."
		t.Errorf(message)
	}
}
