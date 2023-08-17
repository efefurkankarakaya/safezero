package file

import (
	"log"
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
	err := error(nil)
	message := "File overwrote (with fixed size): " + path
	chunkSize := int64(68157440)

	// TODO: Chunk size should be increased for performance but how it can be optimized for every system?
	if size < chunkSize {
		data := utils.GenerateRandomString(size)
		err = os.WriteFile(path, []byte(data), 0644)
	} else {
		// To keep memory usage under control.
		sizeLeft := size
		totalChunkWritten := int64(0)

		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		errorManagement.ReflectError(err)

		// fi, err := f.Stat()
		// errorManagement.ReflectError(err)

		for totalChunkWritten != size {
			if totalChunkWritten > size {
				println("[Chunk Overflow] This should not be happened.")
				// TODO: An event / error log system required.
				os.Exit(1)
			}

			data := utils.GenerateRandomString(chunkSize)

			if _, err := f.Write([]byte(data)); err != nil {
				f.Close() // FileWrite error
				log.Fatal(err)
				break
			} else {
				totalChunkWritten += chunkSize
				sizeLeft -= chunkSize
			}

			if sizeLeft < chunkSize {
				chunkSize = sizeLeft
			}
		}

		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}

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
