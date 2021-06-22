package res

type Result struct {
	Code  int
	Error string
	Data  string
}

func (this *Result) SetResult(code int, msg, data string) *Result {
	this.Code = code
	this.Error = msg
	this.Data = data
	return this
}

func (this *Result) GetResult() *Result {
	return this
}

func ResultInstance() *Result {
	return new(Result)
}
