package test

import (
	"os"
	"path/filepath"
	fileOperations "safezero/internal/file"
	"testing"
)

// $ go test -v ./test

func TestCreatingTestingEnvironment(t *testing.T) {
	// TODO: Refactor here.
	const want bool = true

	var err error
	var message, destination string
	var got bool

	dir, err := os.Getwd()

	if err != nil {
		message = "Error occurred: os.Getwd()"
		t.Errorf(message)
	}

	err = fileOperations.CreateTestingFolder(dir)

	if err != nil {
		println(err.Error())
		message = "Error occurred: CreateTestingFolder()"
		t.Errorf(message)
	}

	err = fileOperations.CopyTestingFilesToTestingFolder(dir)

	if err != nil {
		message = "Error occurred: CopyTestingFilesToTestingFolder"
		t.Errorf(message)
	}

	destination = "temp"
	got = fileOperations.CheckIfPathExists(destination)

	if got != want {
		message = "Failed. '" + destination + "' does not exist."
		t.Errorf(message)
	}

	destination = "temp/lotofen.jpg"
	got = fileOperations.CheckIfPathExists(destination)

	if got != want {
		message = "Failed. '" + destination + "' does not exist."
		t.Errorf(message)
	}
}

func TestFileOverwriting(t *testing.T) {
	const want int64 = 0

	var got int64
	var message string
	var size int64
	var file os.FileInfo

	const testDir string = "temp/"

	filepath.Walk(testDir, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		file, err = os.Stat(path)
		if err != nil {
			return err
		}

		size = file.Size()

		fileOperations.OverwriteFileZeroBytes(path)

		file, err = os.Stat(path)
		if err != nil {
			return err
		}

		size = file.Size()

		got = size

		return err
	})

	if got != want {
		message = "Overwrite test failed."
		t.Errorf(message)
	}
}

func TestFileDeletion(t *testing.T) {
	const want bool = false
	const testDir string = "temp/"

	var got bool
	var message string

	filepath.Walk(testDir, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// TODO: Add remove directory here.
		if !fileInfo.IsDir() {
			fileOperations.RemoveFile(path)
			got = fileOperations.CheckIfPathExists(path)
		}

		return err
	})

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

	err := fileOperations.RemoveTree(testDir)
	if err != nil {
		message = "Error occurred while removing tree."
		t.Errorf(message)
	}

	got = fileOperations.CheckIfPathExists(testDir)

	if got != want {
		message = "Could not removed tree."
		t.Errorf(message)
	}
}
