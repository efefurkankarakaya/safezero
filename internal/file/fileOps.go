package file

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	errorManagement "safezero/internal/error"
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
	errorManagement.PrintSuccess(message, err)
	// if err := os.Mkdir(path, os.ModePerm); err != nil {
	// 	log.Fatal(err)
	// }
	return err
}

func RemoveDirectory(path string) error {
	message := "Folder removed: " + path
	err := os.Remove(path)
	errorManagement.PrintSuccess(message, err)
	return err
}

func RemoveTree(root string) error {
	message := "Tree removed: " + root
	err := os.RemoveAll(root)
	errorManagement.PrintSuccess(message, err)
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

func RemoveFile(path string) error {
	message := "File removed: " + path
	err := os.Remove(path)
	errorManagement.PrintSuccess(message, err)
	return err
}
