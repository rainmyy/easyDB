package bootstrap

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/easydb/library/conf"
	PoolLib "github.com/easydb/library/pool"
)

/**
*app执行入口
 */
type AppServer struct {
	mutex  sync.WaitGroup
	Ctx    context.Context
	Cancel context.CancelFunc
}

const (
	RPCSERVICE = iota
	TCPSERVICE
	READSERVICE
	WRITESERVICE
)

var SERVICELEN = 4
var pool = PoolLib.GetInstance()

/**
*注册执行函数,默认开启rpc服务、tcp服务、数据读服务、数据写服务
 */
func (app *AppServer) Setup() {
	_ = conf.ConfIntance().Init()
	//注册执行函数
	pool := pool.Init(SERVICELEN, SERVICELEN)
	for i := 0; i < SERVICELEN; i++ {
		app.mutex.Add(1)
		go func(num int) {
			defer app.mutex.Done()
			query := PoolLib.QueryInit(strconv.Itoa(num), download, 123, "wwww")
			pool.AddTask(query)
		}(i)
	}
	app.mutex.Wait()
}

func (app *AppServer) Start() {
	pool.Start()
	pool.TaskResult()
}

func download(url int, str string) {
	fmt.Print(str)
	//result := res.ResultInstance().SetResult(200, fmt.Errorf(""), "result")
	//return result
}

func GenInstance() *AppServer {
	return new(AppServer)
}
