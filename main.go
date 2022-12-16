package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "./test-files"

	criticalPaths := []string{
		"",
		"/", "/bin", "/boot", "/cdrom", "/dev", "/etc", "/home", "/lib", "/lost+found", "/media", "/mnt", "/opt", "/proc", "/root", "/run", "/sbin", "/selinux", "/srv", "/sys", "/tmp", "/usr", "/var", "/dev",
		"/Applications", "/Library", "System", "/Users",
		"A:", "B:", "C:", "E:", "D:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P:", "Q:", "R:", "S:", "T:", "U:", "V:", "W:", "X:", "Y:", "Z:",
	}

	for _, criticalPath := range criticalPaths {
		if strings.ToLower(criticalPath) == strings.ToLower(root) {
			fmt.Println(criticalPath)
			os.Exit(1)
		}
	}

	err := filepath.Walk(root, func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		data := []byte("")
		err = ioutil.WriteFile(path, data, 0644)
		fmt.Println("Overwrote: " + path)
		os.Remove(path)
		fmt.Println("Removed: " + path)

		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
