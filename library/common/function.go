package common

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/easydb/library/res"
)

func FuncCall(function interface{}, args ...interface{}) []interface{} {
	var resultList = make([]interface{}, 0)
	result := func(erroMsg string) []interface{} {
		result := res.ResultInstance()
		err := fmt.Errorf(erroMsg)
		result.SetResult(-1, err, "")
		resultList = append(resultList, result)
		return resultList
	}
	value := reflect.ValueOf(function)
	if value.Kind() != reflect.Func {
		return result("params 1 is not a function")
	}
	parameters := make([]reflect.Type, 0, value.Type().NumIn())
	for i := 0; i < value.Type().NumIn(); i++ {
		arg := value.Type().In(i)
		parameters = append(parameters, arg)
	}
	if value.Type().NumIn() != len(args) {
		erroMsg := fmt.Sprintf("argument %d length doesn't equal to provide length %d \n", value.Type().NumIn(), len(args))
		return result(erroMsg)
	}
	outs := make([]reflect.Type, 0, value.Type().NumOut())
	for i := 0; i < value.Type().NumOut(); i++ {
		arg := value.Type().Out(i)
		outs = append(outs, arg)
	}
	var argValues []reflect.Value
	for i := 0; i < len(args); i++ {
		switch parameters[i] {
		case reflect.TypeOf(int(0)):
			argValues = append(argValues, reflect.ValueOf(args[i].(int)))
		case reflect.TypeOf(bool(false)):
			argValues = append(argValues, reflect.ValueOf(args[i].(bool)))
		case reflect.TypeOf(int16(0)):
			argValues = append(argValues, reflect.ValueOf(args[i].(int16)))
		case reflect.TypeOf(string("")):
			argValues = append(argValues, reflect.ValueOf(args[i].(string)))
		default:
			erroMsg := fmt.Sprintf("unsupport type %s[%s] \n", parameters[i].Kind(), parameters[i].Name())
			return result(erroMsg)
		}
	}
	resultValue := value.Call(argValues)
	for i := 0; i < len(resultValue); i++ {
		switch resultValue[i].Type() {
		case reflect.TypeOf(int(0)):
			resultList = append(resultList, resultValue[i].Interface().(int))
		case reflect.TypeOf(string("")):
			resultList = append(resultList, resultValue[i].Interface().(string))
		default:
			resultList = append(resultList, resultValue[i].Interface().(*res.Result))
		}
	}
	if len(outs) > 1 && len(resultList) == 0 {
		return result("result is empty")
	}
	return resultList
}

func FormatResult(resultList []interface{}) *res.Reponse {
	var response = new(res.Reponse)
	for _, reuslt := range resultList {
		if reuslt == nil {
			continue
		}
		switch reuslt.(type) {
		case *res.Result:
			response.Result = reuslt.(*res.Result)
		case int:
			//退出标志
			if reuslt.(int) == -1 {
				return nil
			}
		}
	}
	return response
}

/**
*解析ini格式配置文件
*desc:
*[test]
*    [..params]
*        name:name1
*        key:value
*    [...params]
*        name:name2
*        key:value
 */
func ParserIniContent(content []byte) {
	if content == nil {
		return
	}
	bytesList := [][]byte{}
	hasSlash := false
	bytes := []byte{}
	if content[len(content)-1] != 10 {
		content = append(content, 10)
	}
	for i := 0; i < len(content); i++ {
		value := content[i]
		//出现斜杠过滤
		if value == 47 {
			hasSlash = true
			continue
		}
		if hasSlash {
			if value == 10 {
				hasSlash = false
			}
			continue
		}
		// 通过\n截取长度
		if value != 10 && value != 32 {
			bytes = append(bytes, value)
		} else if len(bytes) > 0 {
			bytesList = append(bytesList, bytes)
			bytes = []byte{}
		}
	}
	if len(bytesList) == 0 {
		return
	}
	//分隔符数组，91:[ 46:. 58:
	var segment = []int{91, 46}
	infunc := InIntSliceSortedFunc(segment)

	for i := 0; i < len(bytesList); i++ {
		bytes := bytesList[i]
		if len(bytes) == 0 {
			continue
		}
		bytesmap := make(map[string]map[string]string)
		for j := 0; j < len(bytes); j++ {
			if j == 0 && infunc(int(bytes[j])) {
				str := Bytes2str(bytes[1:])
				bytesmap[str] = make(map[string]string)
				fmt.Print(bytesmap)
			}
		}
	}

}

func InIntSliceSortedFunc(stack []int) func(int) bool {
	sort.Ints(stack)
	return func(needle int) bool {
		index := sort.SearchInts(stack, needle)
		return index < len(stack) && stack[index] == needle
	}
}
