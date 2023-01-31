package test

import (
	"os"
	"path/filepath"
	fileOperations "safezero/internal/file"
	"testing"
)

func TestFileOverwritingZeroBytes(t *testing.T) {
	const want int64 = 0
	const testDir string = "temp/"

	var got int64
	var message string
	var size int64
	var file os.FileInfo

	// Create the testing files and directories the same directory with test suits.
	dir, _ := os.Getwd()
	fileOperations.CreateTestingFolder(dir)
	fileOperations.CopyTestingFilesToTestingFolder(dir)

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

	// Clean up
	fileOperations.RemoveTree(testDir)

	if got != want {
		message = "Overwrite test failed."
		t.Errorf(message)
	}
}
