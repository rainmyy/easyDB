package strategy

type Stack struct {
	maxNum int
	top    int
	arr    [20]int
}

func (this *Stack) Push() {
	if this.isFull() {

	}
}

func (this *Stack) isFull() bool {
	return this.top+1 >= this.maxNum
}
