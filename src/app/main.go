package main

import (
	"os"
	"io/ioutil"
	"app/lib"
)

func main() {
	file, err := os.OpenFile("/1.json", os.O_RDONLY, os.ModePerm)
	lib.ErrorPut(err)
	defer file.Close()
	
	fileContent, err := ioutil.ReadAll(file)
	lib.FindRequest(string(fileContent), "")
	
	lib.ErrorPut(err)
}
