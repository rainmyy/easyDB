package conf

import (
	"fmt"
	"sync"

	"github.com/rainmyy/easyDB/library/file"
)

/**
* 获取配置信息，默认获取获取default日志
 */

type deafultConf struct {
	m *sync.RWMutex
}

func (conf *deafultConf) Init() *deafultConf {
	confName := "global.conf"
	filepath := "./conf/idc/bj/"

	fileObj := file.FileInstance(confName, filepath)
	result := make(map[string]interface{})
	fmt.Print(result.(type))
	fileObj.Parser(result)
	return conf
}

func ConfIntance() *deafultConf {
	return new(deafultConf)
}
