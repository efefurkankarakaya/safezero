package safedelete

import (
	"io/fs"
	"path/filepath"
	errorOperations "safezero/internal/error"
	fileOperations "safezero/internal/file"
	"time"
)

// Attackers can notify all the files are 0 bytes and can trace clusters by patterns.
// They won't know which clusters and chains belong to what but can be understood.
// It provides more safety than any OS' default deletion algorithm but not safe.
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

// Attackers can notify all the deleted files are the same size with cluster but tracing patterns are not that easy.
// Divides a file into chunks of chains and makes untraceable by patterns or node links.
// It provides average safety.
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
