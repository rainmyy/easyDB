package conf

import (
	"fmt"
	"sync"

	. "github.com/rainmyy/easyDB/library/bind"
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
	confName := "./conf/idc/bj/global.conf"
	fileObj := file.FileInstance(confName)
	err := fileObj.Parser(common.IniType)
	if err != nil {
		print(err.Error())
		return nil
	}
	str := StrigInstance()
	//array := ArrayInterface()
	bindData := DefaultBind(fileObj.GetContent(), str)
	fmt.Print(bindData)
	//bytes, err := json.Marshal(bindData)
	//fmt.Print(string(bytes))
	return conf
}

func ConfIntance() *DeafultConf {
	return new(DeafultConf)
}
