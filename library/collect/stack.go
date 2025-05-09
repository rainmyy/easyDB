package collect

const treVal = 0.75

type Stack struct {
	list *ListObj
}

func StackInstance() *Stack {
	return &Stack{new(ListObj)}
}
func (this *Stack) Len() int {
	return int(this.list.length)
}
func (this *Stack) Peek() interface{} {
	if this.Len() == 0 {
		return nil
	}
	return this.list.head
}

func (this *Stack) Pop() interface{} {
	if this.Len() == 0 {
		return nil
	}
	value := this.list.head
	result := this.list.Delete(0)
	if !result {
		return nil
	}
	return value
}

func (this *Stack) Push(value interface{}) bool {
	if value == nil {
		return false
	}
	result := this.list.Insert(0, value)
	return result
}
