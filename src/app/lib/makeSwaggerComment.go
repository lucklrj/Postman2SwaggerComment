package lib

import (
	"strings"
	"github.com/tidwall/gjson"
)

type LineComent struct {
	Content   string
	IndentNum int
}

type AllComment []SingleComment

type SingleComment struct {
	Body []string
}

/**
 * @SWG\Post(
 *    tags={"手机微信端-线索-带看"},
 *    path="/v1/seller/reservation/guide/bind_guider",
 *    summary="绑定带看人",
 *    @SWG\Parameter(name="salestoken", type="string", required=true, in="query",description="salestoken"),
 *    @SWG\Parameter(name="reservation_id", type="string", required=true, in="query",description="线索id"),
 *    @SWG\Parameter(name="guider_id", type="integer", required=true, in="query",description="带看人id"),
 *    @SWG\Parameter(name="guide_time", type="string", required=true, in="query",description="预计带看时间"),
 *    @SWG\Parameter(name="guide_space_id", type="integer", required=true, in="query",description="预计带看空间id"),
 

 *    @SWG\Response(
 *         response="200",
 *         description="接口响应",
 *         @SWG\Schema(
 *            @SWG\Property( property="code" , type="integer" , example="500",description="填写描述"),
 *            @SWG\Property( property="message" , type="string" , example="该线索当前状态不允许指派带看人",description="填写描述"),
 *         )
 *    )
 * )
 **/
func MakeComment(singeRequest Request) []string {
	ResponseComment = make([]LineComent, 0)
	BodyComment = make([]LineComent, 0)
	
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
		BodyJson2Comemt(gjson.Parse(singeRequest.Body.Content.(string)), blankIndex)
		for _, singleBodyComment := range BodyComment {
			comment = append(comment, blankRepeat(blankIndex+singleBodyComment.IndentNum)+singleBodyComment.Content)
		}
		comment = append(comment, blankRepeat(blankIndex)+")")
	} else if singeRequest.Body.Mode == "formdata" || singeRequest.Body.Mode == "urlencoded" {
		for _, singleBodyParameter := range singeRequest.Body.Content.([]Parameter) {
			singeBodyParameter := "@SWG\\Parameter(name =\"" + singleBodyParameter.Key + "\", type=\"" + singleBodyParameter.Type + "\", required=true, in=\"body\",description=\"" + singleBodyParameter.Description + "\"),"
			comment = append(comment, blankRepeat(blankIndex)+singeBodyParameter)
		}
	}
	comment[len(comment)-1] = comment[len(comment)-1] + ","
	//Response
	comment = append(comment, blankRepeat(blankIndex)+"@SWG\\Response(")
	comment = append(comment, blankRepeat(blankIndex+1)+"response=\"200\",")
	comment = append(comment, blankRepeat(blankIndex+1)+"description=\"接口响应\",")
	
	ResponseJson2Comemt(gjson.Parse(singeRequest.Response), blankIndex)
	for _, singleResponse := range ResponseComment {
		comment = append(comment, blankRepeat(blankIndex+singleResponse.IndentNum)+singleResponse.Content)
	}
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
