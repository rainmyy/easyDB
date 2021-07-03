package res

import "fmt"

type Result struct {
	Code  int
	Error error
	Data  string
}

func (this *Result) SystemError() *Result {
	return this.SetResult(-1, fmt.Errorf("system error"), "")
}

func (this *Result) ErrorParamsResult() *Result {
	return this.SetResult(-2, fmt.Errorf("error params data"), "")
}

func (this *Result) EmptyResult() *Result {
	return this.SetResult(-1, fmt.Errorf("empty data"), "")
}
func (this *Result) SetResult(code int, msg error, data string) *Result {
	this.Code = code
	this.Error = msg
	if data != "" {
		this.Data = data
	}
	return this
}

func (this *Result) GetResult() *Result {
	return this
}

func ResultInstance() *Result {
	return new(Result)
}
