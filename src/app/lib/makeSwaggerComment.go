package lib

import (
	"strings"
)

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
	
	comment := make([]string, 0)
	blankIndex := 0
	
	//请求方式
	comment = append(comment, "@SWG\\"+singeRequest.Method+"(")
	
	blankIndex = blankIndex + 1
	
	//tags
	tags := strings.Split(singeRequest.Name, "/")[0]
	comment = append(comment, blankRepeat(blankIndex)+"tags={\""+tags+"\"},")
	
	//path
	comment = append(comment, blankRepeat(blankIndex)+"path=\""+singeRequest.Path+",")
	
	//summary
	summary := strings.Replace(singeRequest.Name, "/", "-", -1)
	comment = append(comment, blankRepeat(blankIndex)+"summary=\""+summary+",")
	
	//Parameter
	parameterIndex := 0
	queryNum := len(singeRequest.Query)
	
	for parameterIndex < queryNum {
		singeParameter := "@SWG\\Parameter(name =\"" + singeRequest.Query[parameterIndex].Key + "\", type=\"" + singeRequest.Query[parameterIndex].Type + "\", required=true, in=\"query\",description=\"" + singeRequest.Query[parameterIndex].Description + "\"),"
		comment = append(comment, blankRepeat(blankIndex)+singeParameter)
		parameterIndex = parameterIndex + 1
	}
	
	//Response
	
	
	return comment
	
}
func blankRepeat(num int) string {
	return strRepeat(" ", num)
}
func strRepeat(str string, num int) string {
	return strings.Repeat(str, num)
}
