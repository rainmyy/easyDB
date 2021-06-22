package conf

import "sync"

/**
* 获取配置信息，默认获取获取default日志
 */

type DeafultConf struct {
	m *sync.RWMutex
}
