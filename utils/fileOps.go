package utils

import (
	"fmt"
	"io"
	"io/fs"
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

func generateRandomInteger(min int, max int) int {
	rand.Seed(time.Now().UnixMicro())
	return rand.Intn(max-min) + min
}

func getCharacters() []rune {
	var characters []rune = []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'r', 's', 't', 'u', 'v', 'y', 'z', 'q', 'x', 'w',
		'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
		'\'', '"', '!', '^', '+', '%', '&', '/', '(', ')', '=', '?', '_', ',', ';', '.', '*', '-', '<', '>', '#', '$', '{', '}', '[', ']', '|',
	}
	return characters
}

func getRandomCharacter() string {
	var characters []rune = getCharacters()
	var randomIndex int
	var randomCharacter string

	randomIndex = generateRandomInteger(0, len(characters))
	randomCharacter = string(characters[randomIndex])

	return randomCharacter
}

func OverwriteFileFixedSize(path string) error {
	// This size can be determined by the user, default value is default cluster size (4 KB)
	// https://ux.stackexchange.com/questions/13815/files-size-units-kib-vs-kb-vs-kb
	message := "File overwrote (with 4 KB size): " + path
	var fileSize int
	var data string

	const clusterSize int = 4096

	fileSize = clusterSize

	for i := 0; i < fileSize; i++ {
		data += getRandomCharacter()
	}

	err := os.WriteFile(path, []byte(data), 0644)

	PrintSuccess(message, err)
	return err
}

func OverwriteFileRandomSize(path string) error {
	message := "File overwrote (with random size): " + path
	var fileSize int
	var data string

	const minSize int = 4096
	// const maxSize int = 65536

	// Picks a random file size between default cluster size and max size.
	// fileSize = generateRandomInteger(minSize, maxSize)

	// 4 KB (default cluster size) = 1 Disk Access
	// I think that should be faster than normal random length write operation. Especially in HDDs.
	fileSize = minSize * generateRandomInteger(1, 17)

	for i := 0; i < fileSize; i++ {
		data += getRandomCharacter()
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
