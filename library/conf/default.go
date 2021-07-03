package conf

import (
	"sync"

	"github.com/easydb/library/file"
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

	_ = file.FileInstance(confName, filepath)
	//result := []string{}
	//fileObj.Parser(result)
	return conf
}

func ConfIntance() *deafultConf {
	return new(deafultConf)
}
