package utils

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil" // TODO: Deprecated. Use instead os by adding unit tests.
	"os"
	"path/filepath"
)

// TODO: fileController.go?

func CheckIfPathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func createDirectory(path string) error {
	// TODO: Check if path exists.
	message := "Folder created: " + path
	err := os.Mkdir(path, os.ModePerm)
	PrintSuccess(message, err)
	// if err := os.Mkdir(path, os.ModePerm); err != nil {
	// 	log.Fatal(err)
	// }
	return err
}

func RemoveDirectory(path string) error {
	message := "Folder removed: " + path
	err := os.Remove(path)
	PrintSuccess(message, err)
	return err
}

func RemoveTree(root string) error {
	// TODO: Remove tree.
	// https://pkg.go.dev/os#RemoveAll
	message := "Tree removed: " + root
	err := os.RemoveAll(root)
	PrintSuccess(message, err)
	return err
}

func CopyRecursive(source string, destination string) error {
	err := filepath.Walk(source, func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// TODO: Don't forget to add this to other walk functions.
		if !fileInfo.IsDir() {
			_, err = CopyFile(path, destination+"/"+fileInfo.Name())
		}

		return err
	})

	return err
}

func CopyFile(source, destination string) (int64, error) {
	sourceFileStat, err := os.Stat(source)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", source)
	}

	sourceFile, err := os.Open(source)
	if err != nil {
		return 0, err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return 0, err
	}
	defer destinationFile.Close()

	numberOfBytes, err := io.Copy(destinationFile, sourceFile)
	return numberOfBytes, err
}

func OverwriteFile(path string) error {
	message := "File overwrote: " + path

	data := []byte("")
	err := ioutil.WriteFile(path, data, 0644) // TODO: Should I consider provide a dynamic chmod input to user?
	// err := os.WriteFile(path, data, 0644) // TODO: Test here after migrating from ioutil to os

	PrintSuccess(message, err)
	return err
}

func RemoveFile(path string) error {
	message := "File removed: " + path
	err := os.Remove(path)
	PrintSuccess(message, err)
	return err
}
