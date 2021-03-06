package main

import (
	"fmt"

	"github.com/rainmyy/easyDB/bootstrap"
)

/***
* easydb 服务端入库
 */
func main() {
	bootstrap.GenInstance().Setup()
	bootstrap.GenInstance().Start()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
}
