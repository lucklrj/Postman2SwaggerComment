package lib

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

type LineComent struct {
	Content   string
	IndentNum int
}

type AllComment []SingleComment

type SingleComment struct {
	Body []string
}

func MakeTile(host string, base_Path string, version string, title string, description string, contact string) []string {
	comment := make([]string, 0)
	blankIndex := 0

	comment = append(comment, blankRepeat(blankIndex)+"@SWG\\Swagger(")
	blankIndex = blankIndex + 1

	comment = append(comment, blankRepeat(blankIndex)+"schemes={\"http\",\"https\"},")

	comment = append(comment, blankRepeat(blankIndex)+"host=\""+host+"\",")
	comment = append(comment, blankRepeat(blankIndex)+"basePath=\""+base_Path+"\",")
	comment = append(comment, blankRepeat(blankIndex)+"@SWG\\Info(")

	blankIndex = blankIndex + 1

	comment = append(comment, blankRepeat(blankIndex)+"version=\""+version+"\",")
	comment = append(comment, blankRepeat(blankIndex)+"title=\""+title+"\",")
	comment = append(comment, blankRepeat(blankIndex)+"description=\""+description+"\",")
	comment = append(comment, blankRepeat(blankIndex)+"termsOfService=\"\",")
	comment = append(comment, blankRepeat(blankIndex)+"@SWG\\Contact(")
	blankIndex = blankIndex + 1
	comment = append(comment, blankRepeat(blankIndex)+"email=\""+contact+"\"")

	blankIndex = blankIndex - 1
	comment = append(comment, blankRepeat(blankIndex)+")")

	blankIndex = blankIndex - 1
	comment = append(comment, blankRepeat(blankIndex)+")")
	blankIndex = blankIndex - 1

	comment = append(comment, blankRepeat(blankIndex)+")")

	return comment
}
func MakeComment(singeRequest Request) []string {

	comment := make([]string, 0)
	blankIndex := 0

	//请求方式
	comment = append(comment, "@SWG\\"+singeRequest.Method+"(")

	blankIndex = blankIndex + 1

	//tags
	tags := strings.Split(singeRequest.Name, "/")[0]
	comment = append(comment, blankRepeat(blankIndex)+"tags={\""+tags+"\"},")

	//path
	comment = append(comment, blankRepeat(blankIndex)+"path=\""+singeRequest.Path+"\",")

	//summary
	summary := strings.Replace(singeRequest.Name, "/", "-", -1)
	comment = append(comment, blankRepeat(blankIndex)+"summary=\""+summary+"\",")

	//Parameter
	parameterIndex := 0
	queryNum := len(singeRequest.Query)

	for parameterIndex < queryNum {
		singeParameter := "@SWG\\Parameter(name =\"" + singeRequest.Query[parameterIndex].Key + "\", type=\"" + singeRequest.Query[parameterIndex].Type + "\", required=true, in=\"query\",description=\"" + singeRequest.Query[parameterIndex].Description + "\"),"
		comment = append(comment, blankRepeat(blankIndex)+singeParameter)
		parameterIndex = parameterIndex + 1
	}
	//Body
	if singeRequest.Body.Mode == "raw" {
		comment = append(comment, blankRepeat(blankIndex)+"@SWG\\Schema(")

		bodyComment := make([]LineComent, 0)
		bodyComment = Json2Comemt(gjson.Parse(singeRequest.Body.Content.(string)), blankIndex, bodyComment)

		for _, singleBodyComment := range bodyComment {
			comment = append(comment, blankRepeat(blankIndex+singleBodyComment.IndentNum)+singleBodyComment.Content)
		}
		comment = append(comment, blankRepeat(blankIndex)+"),")

	} else if singeRequest.Body.Mode == "formdata" || singeRequest.Body.Mode == "urlencoded" {
		for _, singleBodyParameter := range singeRequest.Body.Content.([]Parameter) {
			singeBodyParameter := "@SWG\\Parameter(name =\"" + singleBodyParameter.Key + "\", type=\"" + singleBodyParameter.Type + "\", required=true, in=\"body\",description=\"" + singleBodyParameter.Description + "\"),"
			comment = append(comment, blankRepeat(blankIndex)+singeBodyParameter)
		}
	}

	//Response
	comment = append(comment, blankRepeat(blankIndex)+"@SWG\\Response(")
	blankIndex = blankIndex + 1
	comment = append(comment, blankRepeat(blankIndex)+"response=\"200\",")
	comment = append(comment, blankRepeat(blankIndex)+"description=\"接口响应\",")
	comment = append(comment, blankRepeat(blankIndex)+"@SWG\\Schema(")

	responseComment := make([]LineComent, 0)

	fmt.Println(singeRequest.Response)
	responseComment = Json2Comemt(gjson.Parse(singeRequest.Response), blankIndex, responseComment)

	for _, singleResponse := range responseComment {
		comment = append(comment, blankRepeat(blankIndex+singleResponse.IndentNum)+singleResponse.Content)
	}
	comment = append(comment, blankRepeat(blankIndex)+")")
	blankIndex = blankIndex - 1
	comment = append(comment, blankRepeat(blankIndex)+")")
	blankIndex = blankIndex - 1
	comment = append(comment, blankRepeat(blankIndex)+")")
	return comment
}

func blankRepeat(num int) string {
	return strRepeat("  ", num)
}
func strRepeat(str string, num int) string {
	return strings.Repeat(str, num)
}
func Json2Comemt(json gjson.Result, level int, responseComment []LineComent) []LineComent {
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
			responseComment = append(responseComment, line)
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
			responseComment = append(responseComment, line)
			break

		case "JSON":
			if value.IsArray() == true { //返回数据为空时，则设置为空字符串
				len := len(value.Array())
				if len == 0 {
					line := LineComent{}
					line.Content = "@SWG\\Property( property=\"data\" , type=\"string\" , example=\"\",description=\"填写描述\"),"
					line.IndentNum = level
					responseComment = append(responseComment, line)

				} else {
					value = value.Array()[0]

					lineStart := LineComent{}
					lineStart.Content = "@SWG\\Property( property=\"" + key.String() + "\" ,type=\"array\","
					lineStart.IndentNum = level
					responseComment = append(responseComment, lineStart)
					responseComment = Json2Comemt(value, level+1, responseComment)

					lineEnd := LineComent{}
					lineEnd.Content = "),"
					lineEnd.IndentNum = level
					responseComment = append(responseComment, lineEnd)
				}
			} else {
				lineStart := LineComent{}
				lineStart.Content = "@SWG\\Property( property=\"" + key.String() + "\" ,type=\"object\","
				lineStart.IndentNum = level
				responseComment = append(responseComment, lineStart)
				responseComment = Json2Comemt(value, level+1, responseComment)

				lineEnd := LineComent{}
				lineEnd.Content = "),"
				lineEnd.IndentNum = level
				responseComment = append(responseComment, lineEnd)
			}
		case "True", "False":
			//case "False":
			line := LineComent{}
			thisValue := value.String()
			thisType := "bool"

			line.Content = "@SWG\\Property( property=\"" + key.String() + "\" , type=\"" + thisType + "\" , example=\"" + thisValue + "\",description=\"填写描述\"),"
			line.IndentNum = level
			responseComment = append(responseComment, line)
			break
		default:
			fmt.Println(value.Type.String())
		}

		return true
	})
	return responseComment
}
