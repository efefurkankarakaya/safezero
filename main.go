package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"safezero/utils"
	"strings"
)

func main() {
	root := "./test/temp"

	criticalPaths := []string{
		"",
		"/", "/bin", "/boot", "/cdrom", "/dev", "/etc", "/home", "/lib", "/lost+found", "/media", "/mnt", "/opt", "/proc", "/root", "/run", "/sbin", "/selinux", "/srv", "/sys", "/tmp", "/usr", "/var", "/dev",
		"/Applications", "/Library", "System", "/Users",
		"A:", "B:", "C:", "E:", "D:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P:", "Q:", "R:", "S:", "T:", "U:", "V:", "W:", "X:", "Y:", "Z:",
		// TODO: Does Windows still allows to use another partition rather than C as a system partition?
		// TODO: Add more Windows-specific critical paths
	}

	for _, criticalPath := range criticalPaths {
		// TODO: strings.equalFold() here (add with unit test)
		if strings.ToLower(criticalPath) == strings.ToLower(root) {
			fmt.Println(criticalPath)
			os.Exit(1) // TODO: Use instead ReflectError but first rewrite this block as a function
		}
	}

	err := filepath.Walk(root, func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			println(path)
			// utils.OverwriteFile(path)
			// utils.RemoveFile(path)
		}

		return nil
	})

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	utils.ReflectError(err)

	os.Exit(0)
}
