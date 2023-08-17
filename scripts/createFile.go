package main

import (
	"os"
	errorOperations "safezero/internal/error"
	"safezero/utils"
)

/* Higher Size = Higher Memory Consumption */
/* Not recommended using with higher size that exceeds your total memory */
func createFileFast(fileName string, size int64) {
	data := utils.GenerateRandomString(size)
	dir, _ := os.Getwd()
	err := os.WriteFile(dir+"/"+fileName, []byte(data), 0644)
	errorOperations.ReflectError(err)
}

/* Higher Size = Lower Size = Fixed Memory Consumption */
// func createFileSafe() {}

func main() {
	createFileFast("128mb.file", 128*1024*1024)
}
