package lib

import (
	"github.com/tidwall/gjson"
	"strings"
	"fmt"
	"os"
)

var BodyComment []LineComent

func BodyJson2Comemt(json gjson.Result, level int) {
	json.ForEach(func(key, value gjson.Result) bool {
		switch value.Type.String() {
		case "Number":
			line := LineComent{}
			
			thisValue := value.String()
			thisType := ""
			if strings.Contains(thisValue, ".") == true {
				thisType = "float"
			} else {
				thisType = "int"
			}
			line.Content = "@SWG\\Property( property=\"" + key.String() + "\" , type=\"" + thisType + "\" , example=\"" + thisValue + "\",description=\"填写描述\"),"
			line.IndentNum = level
			BodyComment = append(BodyComment, line)
			break
		
		case "String":
			line := LineComent{}
			thisValue := value.String()
			thisType := ""
			if thisValue == "true" || thisValue == "false" {
				thisType = "bool"
			} else {
				thisType = "string"
			}
			line.Content = "@SWG\\Property( property=\"" + key.String() + "\" , type=\"" + thisType + "\" , example=\"" + thisValue + "\",description=\"填写描述\"),"
			line.IndentNum = level
			BodyComment = append(BodyComment, line)
			break
		
		case "JSON":
			lineStart := LineComent{}
			lineStart.Content = "@SWG\\Property( property=\"" + key.String() + "\" ,type=\"object\","
			lineStart.IndentNum = level
			BodyComment = append(BodyComment, lineStart)
			if value.IsArray() == true {
				len := len(value.Array())
				if len == 0 {
					fmt.Println(key.String() + "不能为空数组")
					os.Exit(0)
				} else {
					value = value.Array()[0]
				}
			}
			BodyJson2Comemt(value, level+1)
			
			lineEnd := LineComent{}
			lineEnd.Content = "),"
			lineEnd.IndentNum = level
			BodyComment = append(BodyComment, lineEnd)
		}
		//fmt.Println("key", key.String(), reflect.TypeOf(key.String()), value, value.Type.String())
		
		return true
	})
}
