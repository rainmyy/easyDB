package res

type Reponse struct {
	Name  string
	Value interface{}
}

func ReponseIntance() *Reponse {
	return new(Reponse)
}
