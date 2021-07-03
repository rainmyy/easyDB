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
	//mutex         sync.WaitGroup
	RuntineNumber int
	Total         int
	taskQuery     chan *Queue
	taskResult    chan map[string]*res.Reponse
	taskResponse  map[string]*res.Reponse
}

/**
执行队列
*/
type Queue struct {
	Name     string
	result   chan *res.Reponse
	Excel    *ExcelFunc
	CallBack *CallBackFunc
}
type ExcelFunc struct {
	Name     string
	Function interface{}
	Params   []interface{}
}
type CallBackFunc struct {
	name     string
	Function interface{}
	Params   []interface{}
}

func GetInstance() *Pool {
	return new(Pool)
}
func QueryInit(name string, function interface{}, params ...interface{}) *Queue {
	excelFunc := &ExcelFunc{Function: function, Params: params}
	query := &Queue{Name: name,
		Excel:  excelFunc,
		result: make(chan *res.Reponse, 1),
	}
	return query
}

func (q *Queue) CallBackInit(name string, function interface{}, params ...interface{}) *Queue {
	callBackFunc := &CallBackFunc{name: name, Function: function, Params: params}
	q.CallBack = callBackFunc
	return q
}
func (this *Pool) Init(runtineNumber, total int) *Pool {
	this.RuntineNumber = runtineNumber
	this.Total = total
	this.taskQuery = make(chan *Queue, runtineNumber)
	this.taskResult = make(chan map[string]*res.Reponse, runtineNumber)
	this.taskResponse = make(map[string]*res.Reponse)
	return this
}
func (this *Pool) Start() {
	runtineNumber := this.RuntineNumber
	if len(this.taskQuery) != runtineNumber {
		runtineNumber = len(this.taskQuery)
	}
	var mutex sync.WaitGroup
	for i := 0; i < runtineNumber; i++ {
		mutex.Add(1)
		go func(num int) {
			defer mutex.Done()
			task, ok := <-this.taskQuery
			taskName := task.Name
			result := map[string]*res.Reponse{
				taskName: nil,
			}
			response := res.ReponseIntance()
			if !ok {
				res := res.ResultInstance().ErrorParamsResult()
				response.Result = res
				result[taskName] = response
				this.taskResult <- result
				return
			}
			task.excelQuery()
			taskResult, ok := <-task.result
			if !ok {
				res := res.ResultInstance().EmptyResult()
				response.Result = res
				result[taskName] = response
				this.taskResult <- result
				return
			}
			result = map[string]*res.Reponse{
				taskName: taskResult,
			}
			this.taskResult <- result
		}(i)
	}
	mutex.Wait()
	for i := 0; i < runtineNumber; i++ {
		if result, ok := <-this.taskResult; ok {
			for name, value := range result {
				this.taskResponse[name] = value
			}
		}
	}
}

func (this *Pool) TaskResult() map[string]*res.Reponse {
	return this.taskResponse
}

func (this *Pool) Stop() {
	close(this.taskResult)
}

func (this *Pool) AddTask(task *Queue) {
	this.taskQuery <- task
}

/**
* 执行队列
 */
func (qeury *Queue) excelQuery() {
	defer close(qeury.result)
	excelFunc := qeury.Excel.Function
	if excelFunc == nil {
		return
	}
	var requestChannel = make(chan []interface{})
	go func() {
		defer close(requestChannel)
		params := qeury.Excel.Params
		result := FuncCall(excelFunc, params...)
		if result == nil {
			return
		}
		requestChannel <- result
	}()
	result, ok := <-requestChannel
	if !ok {
		return
	}
	response := FormatResult(result)
	if response == nil {
		return
	}
	var callBackChannel = make(chan []interface{})
	go func() {
		defer close(callBackChannel)
		if qeury.CallBack == nil {
			return
		}
		result := FuncCall(qeury.CallBack.Function, qeury.CallBack.Params...)
		if result == nil {
			return
		}
		callBackChannel <- result
	}()
	resultList, ok := <-callBackChannel
	if !ok && response != nil {
		qeury.result <- response
		return
	}
	callBackResponse := FormatResult(resultList).Result
	if callBackResponse != nil {
		response.Callback = callBackResponse
	}
	qeury.result <- response
}
