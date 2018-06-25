package main

import (
	"os"
	"app/lib"
	"io/ioutil"
	"github.com/fatih/color"
)

func main() {
	file, err := os.OpenFile("/Users/mxj/test/test.json", os.O_RDONLY, os.ModePerm)
	lib.ErrorPut(err)
	defer file.Close()
	
	fileContent, err := ioutil.ReadAll(file)
	lib.FindRequest(string(fileContent), "")
	lib.ErrorPut(err)
	requestNum := len(lib.AllRequest)
	
	i := 0
	commentString := ""
	for i < requestNum {
		comment := lib.MakeComment(lib.AllRequest[i])
		commentString = joinComment(commentString, "/**")
		for _, c := range comment {
			commentString = joinComment(commentString, " *"+c)
		}
		
		commentString = joinComment(commentString, " */")
		commentString = joinComment(commentString, "\n\n")
		i = i + 1
	}
	f, err := os.OpenFile("result.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		color.Red(err.Error())
	}
	defer f.Close()
	
	f.WriteString(commentString)
	color.Green("Done!")
}

func joinComment(source, newLine string) string {
	if source != "" {
		return source + "\n" + newLine
	} else {
		return newLine
	}
}
