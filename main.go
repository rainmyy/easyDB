package main

import (
	"github.com/easydb/bootstrap"
)

/***
* easydb 服务端入库
 */
func main() {
	bootstrap.GenInstance().Setup()
	bootstrap.GenInstance().Start()
}
