package bootstrap

import (
	"context"
)

/**
*app执行入口
 */
type AppServer struct {
	Ctx    context.Context
	Cancel context.CancelFunc
}

/**
*注册执行函数
 */
func (app *AppServer) Setup() {

}

func (app *AppServer) Start() {

}
