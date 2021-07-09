package conf

import (
	"bytes"
	"sync"

	"github.com/rainmyy/easyDB/library/bind"
	"github.com/rainmyy/easyDB/library/common"
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
	dataTree, err := fileObj.Parser(common.IniType)
	if err != nil {
		return nil
	}
	var buffer = new(bytes.Buffer)
	buffer.WriteRune(common.LeftRrance)
	bind.BindString(dataTree, buffer)
	buffer.WriteRune(common.RightRrance)
	print(buffer.String())
	return conf
}

func ConfIntance() *deafultConf {
	return new(deafultConf)
}
