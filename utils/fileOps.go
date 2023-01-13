package utils

import (
	"fmt"
	"io"
	"io/fs" // TODO: Deprecated. Use instead os by adding unit tests.
	"math/rand"
	"os"
	"path/filepath"
	"time"
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

func OverwriteFileZeroBytes(path string) error {
	message := "File overwrote: " + path

	data := []byte("")
	err := os.WriteFile(path, data, 0644)

	PrintSuccess(message, err)
	return err
}

// For SSDs
func OverwriteFileRandomSize(path string) error {
	message := "File overwrote (with random size): " + path
	var characterLength int
	var data string
	var randomIndex int

	var characters []rune = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'r', 's', 't', 'u', 'v', 'y', 'z', 'q', 'x', 'w'}
	var numbers []rune = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	var specials []rune = []rune{'\'', '"', '!', '^', '+', '%', '&', '/', '(', ')', '=', '?', '_', ',', ';', '.', '*', '-', '<', '>', '#', '$', '{', '}', '[', ']', '|'}

	rand.Seed(time.Now().UnixMicro())
	const min int = 4095
	const max int = 65535

	characterLength = rand.Intn(max-min+1) + min

	var typeOfCharacter int
	for i := 0; i < characterLength; i++ {
		typeOfCharacter = rand.Intn(3)
		switch typeOfCharacter {
		case 0:
			randomIndex = rand.Intn(len(characters))
			data += string(characters[randomIndex])
		case 1:
			randomIndex = rand.Intn(len(numbers))
			data += string(numbers[randomIndex])
		case 2:
			randomIndex = rand.Intn(len(specials))
			data += string(specials[randomIndex])
		}
	}

	err := os.WriteFile(path, []byte(data), 0644)

	PrintSuccess(message, err)
	return err
}

// For both SSDs and HDDs
func OverwriteFileEqualSize(path string) {}

func RemoveFile(path string) error {
	message := "File removed: " + path
	err := os.Remove(path)
	PrintSuccess(message, err)
	return err
}
