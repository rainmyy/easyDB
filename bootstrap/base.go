package bootstrap

import (
	"context"
	"strconv"

	PoolLib "github.com/easydb/library/pool"
	"github.com/easydb/library/res"
)

/**
*app执行入口
 */
type AppServer struct {
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
var pool *PoolLib.Pool

/**
*注册执行函数,默认开启rpc服务、tcp服务、数据读服务、数据写服务
 */
func (app *AppServer) Setup() {
	pool.Init(SERVICELEN, SERVICELEN)
	for i := 0; i < SERVICELEN; i++ {
		request := res.RequestIntance()
		request.RequestType = 0
		request.Name = strconv.Itoa(i)
		request.Func = func() *res.Result {
			//return download(url)
		}
		pool.AddTask(request)

	}
}

func (app *AppServer) Start() {
	pool.Start()
}
