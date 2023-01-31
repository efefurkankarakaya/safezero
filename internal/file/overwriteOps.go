package file

import (
	"os"
	errorManagement "safezero/internal/error"
	"safezero/utils"
)

func OverwriteFileZeroBytes(path string) error {
	message := "File overwrote: " + path

	data := []byte("")
	err := os.WriteFile(path, data, 0644)

	errorManagement.PrintSuccess(message, err)
	return err
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
		data += utils.GetRandomCharacter()
	}

	err := os.WriteFile(path, []byte(data), 0644)

	errorManagement.PrintSuccess(message, err)
	return err
}

func OverwriteFileRandomSize(path string) error {
	message := "File overwrote (with random size): " + path
	var fileSize int
	var data string

	const minSize int = 4096
	const maxSize int = 65536

	// Picks a random file size between default cluster size and max size.
	fileSize = utils.GenerateRandomInteger(minSize, maxSize)

	// 4 KB (default cluster size) = 1 Disk Access
	// I think that should be faster than normal random length write operation. Especially in HDDs.
	// fileSize = minSize * generateRandomInteger(1, 17)

	for i := 0; i < fileSize; i++ {
		data += utils.GetRandomCharacter()
	}

	err := os.WriteFile(path, []byte(data), 0644)

	errorManagement.PrintSuccess(message, err)
	return err
}

// For both SSDs and HDDs
func OverwriteFileEqualSize(path string) {}
