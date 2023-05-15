package safewipe

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	errorOperations "safezero/internal/error"
	fileOperations "safezero/internal/file"
	"strconv"
)

func WipeSafe(root string) {
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
}

func WipeSafePlus(root string, pass int) {
	err := filepath.Walk(root, func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		file, err := os.Lstat(path)
		if err != nil {
			return err
		}

		size := file.Size()

		// Check if file is regular (non-symlink) and not a directory.
		// TODO: Need to be logged this non-regular files.
		if fileInfo.Mode().IsRegular() && !fileInfo.IsDir() {
			for i := 0; i < pass; i++ {
				fileOperations.OverwriteFileFixedSize(path, size)
			}
			fileOperations.RemoveFile(path)
		} else if !fileInfo.IsDir() {
			println(fmt.Sprintf(`*
File: %s, 
Type: %s,
Mode: %s,
Is Regular?: %s,
Is Dir?: %s
`,
				path,
				file.Mode().Type().String(),
				file.Mode().String(),
				strconv.FormatBool((file.Mode().IsRegular())),
				strconv.FormatBool(file.Mode().IsDir())))
			// println("File: " + path)
			// println("Type: " + (file.Mode().Type().String()))
			// println("Mode: " + file.Mode().String())
			// println("Is regular?: " + strconv.FormatBool((file.Mode().IsRegular())))
			// println("Is Dir?: " + strconv.FormatBool(file.Mode().IsDir()))
		}

		return err
	})
	errorOperations.ReflectError(err)

	err = fileOperations.RemoveTree(root)
	errorOperations.ReflectError(err)
}
