package safedelete

import (
	"io/fs"
	"path/filepath"
	errorOperations "safezero/internal/error"
	fileOperations "safezero/internal/file"
	"time"
)

func DeleteFast(root string) {
	err := filepath.Walk(root, func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			fileOperations.OverwriteFileZeroBytes(path)
			fileOperations.RemoveFile(path)
		}

		return err
	})
	errorOperations.ReflectError(err)

	err = fileOperations.RemoveTree(root)
	errorOperations.ReflectError(err)
}

func DeleteSecure(root string) {
	// TODO: Remove time
	start := time.Now()

	// TODO: 1) User can specify cluster size
	// TODO: 2) Cluster size can be obtained from partition defaults
	// TODO: 3) If not passed as input by user or could not get from partition defaults, use 4 KB
	const clusterSize = 4096

	err := filepath.Walk(root, func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			fileOperations.OverwriteFileFixedSize(path, clusterSize)
			fileOperations.RemoveFile(path)
		}

		return err
	})
	errorOperations.ReflectError(err)

	err = fileOperations.RemoveTree(root)
	errorOperations.ReflectError(err)

	println(time.Since(start))
}

func DeleteSecurePlus(root string) {
	err := filepath.Walk(root, func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			fileOperations.OverwriteFileRandomSize(path)
			fileOperations.OverwriteFileZeroBytes(path)
			fileOperations.RemoveFile(path)
		}

		return err
	})
	errorOperations.ReflectError(err)

	err = fileOperations.RemoveTree(root)
	errorOperations.ReflectError(err)
}
