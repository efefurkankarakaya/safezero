package file

func CreateTestingFolder(root string) error {
	// TODO: Consider os.TempDir()
	const folderName string = "temp"
	testingFolder := root + "/" + folderName

	isExisted := CheckIfPathExists(testingFolder)

	if !isExisted {
		err := createDirectory(testingFolder)
		return err
	}

	return nil
}

func CopyTestingFilesToTestingFolder(root string) error {
	var folderName string

	folderName = "files"
	testingFiles := root + "/" + folderName

	folderName = "temp"
	testingFolder := root + "/" + folderName
	isTestingFolderExisted := CheckIfPathExists(testingFolder)

	if isTestingFolderExisted {
		err := CopyRecursive(testingFiles, testingFolder)

		// ReflectError(err)
		return err
	}
	return nil
}
