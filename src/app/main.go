package main

import (
	"os"
	"app/lib"
	"io/ioutil"
	"fmt"
)

func main() {
	file, err := os.OpenFile("/Users/mxj/1.json", os.O_RDONLY, os.ModePerm)
	lib.ErrorPut(err)
	defer file.Close()
	
	fileContent, err := ioutil.ReadAll(file)
	lib.FindRequest(string(fileContent), "")
	lib.ErrorPut(err)
	
	requestNum := len(lib.AllRequest)
	
	i := 0
	for i < requestNum {
		comment := lib.MakeComment(lib.AllRequest[i])
		fmt.Println(comment)
		i = i + 1
	}
}
