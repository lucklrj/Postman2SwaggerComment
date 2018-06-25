package main

import (
	"os"
	"app/lib"
	"io/ioutil"
	"github.com/fatih/color"
	"flag"
)

var (
	inputFile  = flag.String("source", "", "Postman 导出的json文件")
	outputFile = flag.String("output", "", "产生的注释文件")
)

func init() {
	flag.Parse()
	if *inputFile == "" {
		color.Red("缺少参数：source")
		os.Exit(0)
	}
	if *outputFile == "" {
		*outputFile = "result.txt"
	}
}
func main() {
	file, err := os.OpenFile(*inputFile, os.O_RDONLY, os.ModePerm)
	lib.ErrorPut(err)
	defer file.Close()
	
	fileContent, err := ioutil.ReadAll(file)
	lib.FindRequest(string(fileContent), "")
	lib.ErrorPut(err)
	requestNum := len(lib.AllRequest)
	if requestNum == 0 {
		color.Red("没有找到request数据")
		os.Exit(0)
	}
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
	f, err := os.OpenFile(*outputFile, os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()
	if err != nil {
		color.Red(err.Error())
	} else {
		f.WriteString(commentString)
		color.Green("Done!")
	}
}

func joinComment(source, newLine string) string {
	if source != "" {
		return source + "\n" + newLine
	} else {
		return newLine
	}
}
