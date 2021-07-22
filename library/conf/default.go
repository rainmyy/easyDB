package conf

import (
	"sync"

	"github.com/rainmyy/easyDB/library/bind"
	"github.com/rainmyy/easyDB/library/common"
	"github.com/rainmyy/easyDB/library/file"
)

/**
* 获取配置信息，默认获取获取default日志,配置文件bind级别分为must,should,must是必须参数，should是非必须参数，不配置则不判断bind
 */
type DeafultConf struct {
	m     *sync.RWMutex
	Key   string         `bind:"must"`
	Name  string         `bind:"should"`
	Child []*DeafultConf `bind:"must"`
}

func (conf *DeafultConf) Init() *DeafultConf {
	confName := "global.conf"
	filepath := "./conf/idc/bj/"
	fileObj := file.FileInstance(confName, filepath, 1)
	err := fileObj.Parser(common.IniType)
	if err != nil {
		print(err.Error())
		return nil
	}
	//dataTree := common.Bytes2TreeStruct(fileObj.GetContent())
	str := bind.StrigInstance()
	bind.DefaultBind(fileObj.GetContent(), str)
	return conf
}

func ConfIntance() *DeafultConf {
	return new(DeafultConf)
}
