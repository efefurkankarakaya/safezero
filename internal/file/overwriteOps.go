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

func OverwriteFileFixedSize(path string, size int64) error {
	// This size can be determined by the user, default value is default cluster size (4 KB)
	// https://ux.stackexchange.com/questions/13815/files-size-units-kib-vs-kb-vs-kb
	message := "File overwrote (with fixed size): " + path

	data := utils.GenerateRandomString(size)

	err := os.WriteFile(path, []byte(data), 0644)

	errorManagement.PrintSuccess(message, err)
	return err
}

func OverwriteFileRandomSize(path string) error {
	message := "File overwrote (with random size): " + path
	var fileSize int64
	var data string

	const minSize int64 = 4096     // 4 * 1024
	const maxSize int64 = 68157440 // (65 * 1024 * 1024)

	// Picks a random file size between default cluster size and max size.
	fileSize = utils.GenerateRandomInteger(minSize, maxSize)

	// 4 KB (default cluster size) = 1 Disk Access
	// I think that could be faster than normal random length write operation in some situations. Especially in HDDs.
	// fileSize = minSize * generateRandomInteger(1, 17)

	data = utils.GenerateRandomString(fileSize)

	err := os.WriteFile(path, []byte(data), 0644)

	errorManagement.PrintSuccess(message, err)
	return err
}

// For both SSDs and HDDs
func OverwriteFileEqualSize(path string) {}
