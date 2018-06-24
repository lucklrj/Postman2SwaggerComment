package main

import (
	"os"
	"app/lib"
	"io/ioutil"
	"github.com/tidwall/gjson"
	"fmt"
	"strings"
)

func main() {
	file, err := os.OpenFile("/Users/mxj/test/content.json", os.O_RDONLY, os.ModePerm)
	lib.ErrorPut(err)
	defer file.Close()
	
	fileContent, err := ioutil.ReadAll(file)
	lib.ErrorPut(err)
	
	fileContentString := string(fileContent)
	
	json := gjson.Parse(fileContentString)
	IndentNum := 0
	lib.BodyJson2Comemt(json, IndentNum)
	
	len := len(lib.BodyComment)
	
	i := 0
	for i < len {
		fmt.Println(strings.Repeat("    ", lib.BodyComment[i].IndentNum) + lib.BodyComment[i].Content)
		i = i + 1
	}
}
