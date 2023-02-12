package safewipe

import (
	"io/fs"
	"os"
	"path/filepath"
	errorOperations "safezero/internal/error"
	fileOperations "safezero/internal/file"
	"time"
)

func WipeSafe(root string) {
	// TODO: Remove time
	start := time.Now()

	err := filepath.Walk(root, func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		file, err := os.Stat(path)
		if err != nil {
			return err
		}

		size := file.Size()

		if !fileInfo.IsDir() {
			fileOperations.OverwriteFileFixedSize(path, size)
			fileOperations.RemoveFile(path)
		}

		return err
	})
	errorOperations.ReflectError(err)

	err = fileOperations.RemoveTree(root)
	errorOperations.ReflectError(err)

	println(time.Since(start))
}

func WipeSafePlus(root string, pass int) {
	err := filepath.Walk(root, func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		file, err := os.Stat(path)
		if err != nil {
			return err
		}

		size := file.Size()

		if !fileInfo.IsDir() {
			for i := 0; i < pass; i++ {
				fileOperations.OverwriteFileFixedSize(path, size)
			}
			fileOperations.RemoveFile(path)
		}

		return err
	})
	errorOperations.ReflectError(err)

	err = fileOperations.RemoveTree(root)
	errorOperations.ReflectError(err)
}
