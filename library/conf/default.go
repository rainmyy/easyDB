package conf

import (
	"sync"

	"github.com/rainmyy/easyDB/library/bind"
	"github.com/rainmyy/easyDB/library/common"
	"github.com/rainmyy/easyDB/library/file"
)

/**
* 获取配置信息，默认获取获取default日志
 */

type DeafultConf struct {
	m    *sync.RWMutex
	Test []*DeafultConf `bind:"test"`
}

type Params struct {
	Name string `bind:"name"`
	Key  string `bind:"key"`
}

type Count struct {
	Value string `bind:value`
}

func (conf *DeafultConf) Init() *DeafultConf {
	confName := "global.conf"
	filepath := "./conf/idc/bj/"

	fileObj := file.FileInstance(confName, filepath)
	dataTree, err := fileObj.Parser(common.IniType)
	if err != nil {
		return nil
	}
	bind.DefaultBindMap(dataTree)
	return conf
}

func ConfIntance() *deafultConf {
	return new(deafultConf)
}
