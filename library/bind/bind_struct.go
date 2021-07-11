package bind

import (
	"fmt"
	"reflect"

	"github.com/rainmyy/easyDB/library/strategy"
)

func DefaultBindStruct(treeList []*strategy.TreeStruct, obj interface{}) []interface{} {
	st := reflect.TypeOf(obj)
	getBindParams(st)
	return nil
}

func getBindParams(obj interface{}) {
	t := reflect.TypeOf(obj).Elem()
	field := t.Field(0)
	fmt.Print(field)
	//params := make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		//print(obj)
		fmt.Print("------>", t.Field(i).Tag)
		//params[t.Field(i).Tag.Get("bind")] = t.Field(i)
	}
}
