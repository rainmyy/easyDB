package res

type Reponse struct {
	Name     string
	Result   *Result
	Callback *Result
	res      Result
}

func ReponseIntance() *Reponse {
	return new(Reponse)
}
