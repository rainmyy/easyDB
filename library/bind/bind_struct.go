package bind

import (
	"log"
	"reflect"
	"strings"

	"github.com/rainmyy/easyDB/library/strategy"
)

func DefaultBindStruct(treeList []*strategy.TreeStruct, obj interface{}) []interface{} {
	tagsMapList := getBindParams(obj)
	if len(tagsMapList) == 0 {
		return nil
	}
	robj := reflect.ValueOf(obj)
	if robj.Kind() != reflect.Ptr || robj.IsNil() {
		return nil
	}
	var ginBindStruct func()
	ginBindStruct = func() {

	}
	ginBindStruct()
	return nil
}

func getBindParams(obj interface{}) []map[string][]string {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Print("error: type is not struct")
		return nil
	}
	//result := make(map[string]string)
	tagsMapList := make([]map[string][]string, 0)
	for i := 0; i < t.NumField(); i++ {
		//print(obj)
		fieldame := t.Field(i).Name
		fieldTag := t.Field(i).Tag
		newFieldTag := strings.Replace(string(fieldTag), " ", "", -1)
		tags := strings.Split(newFieldTag, ":")
		tagsMap := make(map[string][]string)
		if len(tags) == 0 {
			continue
		}
		emptyTag := false
		for _, tag := range tags {
			if tag == "" {
				emptyTag = true
			}
		}
		//非bind操作tag不执行
		if emptyTag || tags[0] != "bind" {
			continue
		}
		if len(tags) > 1 {
			tags[1] = tags[1][1 : len(tags[1])-1]
		}
		tagsMap[fieldame] = tags
		tagsMapList = append(tagsMapList, tagsMap)
	}
	return tagsMapList
}
