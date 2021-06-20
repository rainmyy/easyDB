package pool

import "fmt"

type Pool struct {
	Queue         chan func() error
	RuntineNumber int
	Total         int

	Result         chan error
	FinishCallback func()
}

func (this *Pool) Init(runtineNumber, total int) {
	this.RuntineNumber = runtineNumber
	this.Total = total
	this.Queue = make(chan func() error, total)
	this.Result = make(chan error, total)
}

func (this *Pool) Start() {
	for i := 0; i < this.RuntineNumber; i++ {
		go func() {
			for {
				task, ok := <-this.Queue
				if !ok {
					break
				}
				err := task()
				this.Result <- err
			}
		}()
	}

	for j := 0; j < this.Total; j++ {
		res, ok := <-this.Result
		if !ok {
			break
		}
		if res != nil {
			fmt.Println(res)
		}
	}

	if this.FinishCallback != nil {
		this.FinishCallback()
	}
}

func (this *Pool) Stop() {
	close(this.Queue)
	close(this.Result)
}

func (this *Pool) AddTask(task func() error) {
	this.Queue <- task
}

func (this *Pool) SetFinishCallback(callback func()) {
	this.FinishCallback = callback
}
