package main

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	errorOperations "safezero/internal/error"
	fileOperations "safezero/internal/file"
	"strings"
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
	err := filepath.Walk(root, func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			fileOperations.OverwriteFileFixedSize(path)
			fileOperations.RemoveFile(path)
		}

		return err
	})
	errorOperations.ReflectError(err)

	err = fileOperations.RemoveTree(root)
	errorOperations.ReflectError(err)
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

func Erase() {}

func main() {
	// TODO: Should I consider cryptographic seed instead of time seed?
	rand.Seed(time.Now().UnixMicro())

	root := "/Users/efefurkankarakaya/Downloads/contemporary"

	criticalPaths := []string{
		"",
		"/", "/bin", "/boot", "/cdrom", "/dev", "/etc", "/home", "/lib", "/lost+found", "/media", "/mnt", "/opt", "/proc", "/root", "/run", "/sbin", "/selinux", "/srv", "/sys", "/tmp", "/usr", "/var", "/dev",
		"/Applications", "/Library", "System", "/Users",
		"A:", "B:", "C:", "E:", "D:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P:", "Q:", "R:", "S:", "T:", "U:", "V:", "W:", "X:", "Y:", "Z:",
		// TODO: Does Windows still allow to use another partition rather than C as a system partition?
		// TODO: Add more Windows-specific critical paths.
	}

	for _, criticalPath := range criticalPaths {
		if strings.EqualFold(criticalPath, root) {
			fmt.Println(criticalPath)
			os.Exit(1) // TODO: Use instead ReflectError but first rewrite this block as a function.
		}
	}

	// DeleteSecure(root)
	// DeleteFast(root)
	DeleteSecurePlus(root)

	os.Exit(0)
}
