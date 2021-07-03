package res

type Reponse struct {
	Name     string
	Result   *Result
	Callback *Result
	//res      Result
}

func ReponseIntance() *Reponse {
	return new(Reponse)
}

func (this *Reponse) ErrorMsg(tag, errcode int, errmsg error) {
	result := ResultInstance().SetResult(errcode, errmsg, "")
	if tag == 1 {
		this.Result = result
	} else if tag == 2 {
		this.Callback = result
	}
}
