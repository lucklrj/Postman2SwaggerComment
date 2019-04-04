package main

import (
	"app/lib"
	"flag"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

var (
	inputFile  = flag.String("source", "", "Postman 导出的json文件")
	outputFile = flag.String("output", "", "产生的注释文件")

	host        = flag.String("host", "Project host", "项目地址,例如：api.xx.com")
	basePath    = flag.String("base_path", "/", "项目地址,例如 /")
	title       = flag.String("title", "Project Api", "项目名称")
	description = flag.String("description", "Project description", "项目描述")
	version     = flag.String("version", "", "项目版本号")
	contact     = flag.String("contact", "", "联系方式")
)

func init() {
	flag.Parse()
	if *inputFile == "" {
		color.Red("缺少参数：source")
		os.Exit(0)
	}
	if *outputFile == "" {
		*outputFile = "result.php"
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
	commentString := "<?php\n/**"
	comment := lib.MakeTile(*host, *basePath, *version, *title, *description, *contact)
	for _, c := range comment {
		commentString = joinComment(commentString, " *"+c)
	}
	commentString = joinComment(commentString, " */\n\n")

	for i < requestNum {
		comment = lib.MakeComment(lib.AllRequest[i])
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
