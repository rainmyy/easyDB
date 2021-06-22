package conf

import (
	"sync"

	"github.com/Unknwon/goconfig"
)

/**
* 获取配置信息，默认获取获取default日志
 */

type DeafultConf struct {
	m *sync.RWMutex
}

func (conf *DeafultConf) Init() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
}
