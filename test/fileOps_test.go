package test

import (
	"os"
	"testing"
	"zeror/utils"
)

// $ go test -v ./test

func TestCreatingTestingEnvironment(t *testing.T) {
	// TODO: Refactor here.
	const want bool = true

	var err error
	var message, destination string
	var got bool

	dir, err := os.Getwd()

	if err != nil {
		message = "Error occurred: os.Getwd()"
		t.Errorf(message)
	}

	err = utils.CreateTestingFolder(dir)

	if err != nil {
		println(err.Error())
		message = "Error occurred: CreateTestingFolder()"
		t.Errorf(message)
	}

	err = utils.CopyTestingFilesToTestingFolder(dir)

	if err != nil {
		message = "Error occurred: CopyTestingFilesToTestingFolder"
		t.Errorf(message)
	}

	destination = "temp"
	got = utils.CheckIfPathExists(destination)

	if got != want {
		message = "Failed. '" + destination + "' does not exist."
		t.Errorf(message)
	}

	destination = "temp/lotofen.jpg"
	got = utils.CheckIfPathExists(destination)

	if got != want {
		message = "Failed. '" + destination + "' does not exist."
		t.Errorf(message)
	}
}

func TestFileOverwriting(t *testing.T) {

}

func TestFileDeletion(t *testing.T) {

}
