package pool

import (
	"sync"

	res "github.com/easydb/library/res"
)

const (
	defaultRuntineNumber = 10
	defailtTotal         = 10
)

type Pool struct {
	mutex          sync.WaitGroup
	Queue          chan *res.Request
	RuntineNumber  int
	Total          int
	result         chan map[string]*res.Result
	taskResponse   []*res.Reponse
	FinishCallback map[string]func()
}

func (this *Pool) Init(runtineNumber, total int) {
	this.RuntineNumber = runtineNumber
	this.Total = total
	this.Queue = make(chan *res.Request, runtineNumber)
	this.result = make(chan map[string]*res.Result, runtineNumber)
}

func (this *Pool) Start() {
	runtineNumber := this.RuntineNumber
	if runtineNumber <= 0 {
		runtineNumber = defaultRuntineNumber
	}
	for i := 0; i < runtineNumber; i++ {
		this.mutex.Add(1)
		go func(num int) {
			defer this.mutex.Done()
			task, ok := <-this.Queue
			if !ok {
				return
			}
			taskName := task.Name
			taskResult := task.Func()
			result := map[string]*res.Result{
				taskName: taskResult,
			}
			this.result <- result
		}(i)
	}
	this.mutex.Wait()
	this.taskResponse = []*res.Reponse{}
	for j := 0; j < this.RuntineNumber; j++ {
		result, ok := <-this.result
		if !ok {
			break
		}
		response := res.ReponseIntance()
		for key, value := range result {
			response.Name = key
			response.Value = value
			this.taskResponse = append(this.taskResponse, response)
		}
	}
	//执行回调
	if this.FinishCallback != nil {
		for _, function := range this.FinishCallback {
			function()
		}
	}
}

func (this *Pool) TaskResult() []*res.Reponse {
	return this.taskResponse
}

func (this *Pool) Stop() {
	close(this.Queue)
	close(this.result)
}

func (this *Pool) AddTask(task *res.Request) {
	this.Queue <- task
}

func (this *Pool) SetFinishCallback(name string, callback func()) {
	this.FinishCallback[name] = callback
}
