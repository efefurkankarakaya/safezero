package main

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"safezero/utils"
	"strings"
	"time"
)

func DeleteFast(root string) {
	err := filepath.Walk(root, func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			utils.OverwriteFileZeroBytes(path)
			utils.RemoveFile(path)
		}

		return err
	})
	utils.ReflectError(err)

	err = utils.RemoveTree(root)
	utils.ReflectError(err)
}

func DeleteSecure(root string) {
	err := filepath.Walk(root, func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			utils.OverwriteFileFixedSize(path)
			utils.RemoveFile(path)
		}

		return err
	})
	utils.ReflectError(err)

	err = utils.RemoveTree(root)
	utils.ReflectError(err)
}

func DeleteSecurePlus(root string) {
	err := filepath.Walk(root, func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			utils.OverwriteFileRandomSize(path)
			utils.OverwriteFileZeroBytes(path)
			utils.RemoveFile(path)
		}

		return err
	})
	utils.ReflectError(err)

	err = utils.RemoveTree(root)
	utils.ReflectError(err)
}

func Erase() {}

func main() {
	rand.Seed(time.Now().UnixMicro())

	root := ""

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

	DeleteSecure(root)
	// DeleteFast(root)
	// DeleteSecurePlus(root)

	os.Exit(0)
}
