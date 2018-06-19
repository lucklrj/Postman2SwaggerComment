package lib

import (
	"github.com/tidwall/gjson"
	"reflect"
)

type Parameter struct {
	Key         string
	Value       string
	Description string
	Type        string
}

type Body struct {
	Mode    string
	Content string
}

type Request struct {
	Name     string
	Method   string
	Host     string
	Path     string
	Body     Body
	Response string
	Query    []Parameter
}

var AllRequest []Request

func FindRequest(str string, ParentName string) {
	json := gjson.Parse(str)
	item := json.Get("item")
	itemExists := item.Exists()
	if itemExists == true {
		name := json.Get("name").String()
		item.ForEach(func(key, value gjson.Result) bool {
			FindRequest(value.String(), setName(ParentName, name))
			return true
		})
	} else { //
		ParseRequest(json, ParentName)
	}
}

func ParseRequest(body gjson.Result, ParentName string) {
	name := ParentName + "/" + body.Get("name").String()
	method := body.Get("request.method").String()
	host := body.Get("request.url.protocol").String() + "://" + joinArrayFromInterface(body.Get("request.url.host").Value(), ".")
	path := joinArrayFromInterface(body.Get("request.url.path").Value(), "/")
	query := parseQuery(body.Get("request.url.query").Value().([]interface{}))
	Response := body.Get("response.0.body").String()
	
	bodyRequest := Body{}
	
	bodyMode := body.Get("request.body.mode")
	if bodyMode.Exists() {
		bodyRequest.Mode = bodyMode.String()
		if bodyMode.String() == "raw" {
			bodyRequest.Content = body.Get("request.body.raw").String()
		} else if bodyMode.String() == "formdata" {
			bodyRequest.Content = body.Get("request.body.formdata").String()
		} else if bodyMode.String() == "urlencoded" {
			bodyRequest.Content = body.Get("request.body.urlencoded").String()
		}
	}
	AllRequest = append(AllRequest, Request{Name: name, Method: method, Host: host, Path: path, Query: query, Response: Response, Body: bodyRequest})
	
}
func joinArrayFromInterface(data interface{}, sign string) string {
	returnString := ""
	for index, value := range data.([]interface{}) {
		if index == 0 {
			returnString = value.(string)
		} else {
			returnString = returnString + sign + value.(string)
		}
	}
	
	return returnString
}

func parseQuery(data []interface{}) (formatQuery []Parameter) {
	var singleParameter Parameter
	
	for _, singlePoint := range data {
		
		singeData := singlePoint.(map[string]interface{})
		singleParameter = Parameter{}
		
		for key, value := range singeData {
			
			singleValue := ""
			singleType := ""
			if reflect.TypeOf(value) == nil {
				singleValue = ""
				singleType = "string"
			} else if reflect.TypeOf(value).String() == "string" {
				singleValue = value.(string)
				singleType = "string"
			} else if reflect.TypeOf(value).String() == "bool" {
				if value.(bool) == true {
					singleValue = "true"
				} else {
					singleValue = "false"
				}
				singleType = "bool"
			}
			
			switch key {
			case "key":
				singleParameter.Key = singleValue
			case "value":
				singleParameter.Value = singleValue
			case "description":
				singleParameter.Description = singleValue
			}
			singleParameter.Type = singleType
			
		}
		formatQuery = append(formatQuery, singleParameter)
	}
	return formatQuery
}
func setName(prefixName string, name string) string {
	if prefixName == "" {
		return name
	} else {
		return prefixName + "/" + name
	}
}
