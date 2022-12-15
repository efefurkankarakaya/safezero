package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	root := ""
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
