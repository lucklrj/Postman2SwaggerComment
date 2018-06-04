package lib

import (
	"github.com/tidwall/gjson"
	"reflect"
)

type Request struct {
	Name     string
	Method   string
	Host     string
	Path     string
	Response string
	Query    []map[string]string
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
	AllRequest = append(AllRequest, Request{Name: name, Method: method, Host: host, Path: path, Query: query, Response: Response})
	
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

func parseQuery(data []interface{}) []map[string]string {
	returnMap := make([] map[string]string, 0)
	for _, singlePoint := range data {
		singleMap := make(map[string]string)
		singeData := singlePoint.(map[string]interface{})
		for key, value := range singeData {
			if reflect.TypeOf(value) == nil {
				continue
			}
			if reflect.TypeOf(value).String() == "string" {
				singleMap[key] = value.(string)
			}
			if reflect.TypeOf(value).String() == "bool" {
				if value.(bool) == true {
					singleMap[key] = "true"
				} else {
					singleMap[key] = "false"
				}
				
			}
		}
		returnMap = append(returnMap, singleMap)
	}
	return returnMap
}
func setName(prefixName string, name string) string {
	if prefixName == "" {
		return name
	} else {
		return prefixName + "/" + name
	}
}
