package bind

import (
	"log"
	"reflect"

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
	var ginBindStruct func(treeList []*strategy.TreeStruct)
	ginBindStruct = func(treeList []*strategy.TreeStruct) {
		if len(treeList) == 0 {
			return
		}
	}
	ginBindStruct(treeList)
	return nil
}

func getBindParams(obj interface{}) []map[string]string {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Print("error: type is not struct")
		return nil
	}
	tagsMapList := make([]map[string]string, 0)
	for i := 0; i < t.NumField(); i++ {
		tagMap := make(map[string]string)
		fieldame := t.Field(i).Name
		fieldTag := t.Field(i).Tag
		tagName := fieldTag.Get("bind")
		if tagName == "" {
			continue
		}
		tagMap[fieldame] = tagName
		tagsMapList = append(tagsMapList, tagMap)
	}
	return tagsMapList
}
