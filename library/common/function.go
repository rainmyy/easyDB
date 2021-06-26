package common

import (
	"fmt"
	"reflect"

	"github.com/easydb/library/res"
)

func FuncCall(function interface{}, args ...interface{}) []interface{} {
	value := reflect.ValueOf(function)
	if value.Kind() != reflect.Func {
		return nil
	}
	parameters := make([]reflect.Type, 0, value.Type().NumIn())
	for i := 0; i < value.Type().NumIn(); i++ {
		arg := value.Type().In(i)
		//fmt.Printf("argument %d is %s[%s] type \n", i, arg.Kind(), arg.Name())
		parameters = append(parameters, arg)
	}
	if value.Type().NumIn() != len(args) {
		fmt.Printf("argument %d length doesn't equal to provide length %d \n", value.Type().NumIn(), len(args))
		return nil
	}
	outs := make([]reflect.Type, 0, value.Type().NumOut())
	for i := 0; i < value.Type().NumOut(); i++ {
		arg := value.Type().Out(i)
		//fmt.Printf("out %d is %s[%s] type \n", i, arg.Kind(), arg.Name())
		outs = append(outs, arg)
	}
	if value.Type().NumOut() < 1 {
		fmt.Println("outs length must greater than 0")
		return nil
	}
	var argValues []reflect.Value
	for i := 0; i < len(args); i++ {
		switch parameters[i] {
		case reflect.TypeOf(int(0)):
			argValues = append(argValues, reflect.ValueOf(args[i].(int)))
		case reflect.TypeOf(string("")):
			argValues = append(argValues, reflect.ValueOf(args[i]))
		default:
			fmt.Printf("unsupport type %s[%s] \n", parameters[i].Kind(), parameters[i].Name())
			return nil
		}
	}
	resultValue := value.Call(argValues)
	var resultList []interface{}
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
	if len(resultList) == 0 {
		return nil
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
