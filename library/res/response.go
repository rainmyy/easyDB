package res

type Reponse struct {
	Name     string
	Result   *Result
	callback Result
}

func ReponseIntance() *Reponse {
	return new(Reponse)
}
