package main

import (
	"fmt"
	"strconv"

	PoolLib "github.com/easydb/library/pool"
	res "github.com/easydb/library/res"
)

/***
* easydb 服务端入库
 */
func main() {

	urls := []string{
		"http://dlsw.baidu.com/sw-search-sp/soft/44/17448/Baidusd_Setup_4.2.0.7666.1436769697.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/3a/12350/QQ_V7.4.15197.0_setup.1436951158.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/9d/14744/ChromeStandalone_V43.0.2357.134_Setup.1436927123.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/6f/15752/iTunes_V12.2.1.16_Setup.1436855012.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/70/17456/BaiduAn_Setup_5.0.0.6747.1435912002.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/40/12856/QIYImedia_1_06_v4.0.0.32.1437470004.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/42/37473/BaiduSoftMgr_Setup_7.0.0.1274.1436770136.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/49/16988/YoudaoNote_V4.1.0.300_setup.1429669613.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/55/11339/bdbrowserSetup-7.6.100.2089-1212_11000003.1437029629.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/53/21734/91zhushoupc_Windows_V5.7.0.1633.1436844901.exe",
	}

	pool := new(PoolLib.Pool)
	pool.Init(len(urls), len(urls))

	for i := range urls {
		url := urls[i]
		request := res.RequestIntance()
		request.RequestType = 0
		request.Name = "TAST" + strconv.Itoa(i)
		request.Func = func() *res.Result {
			return download(url)
		}
		pool.AddTask(request)
	}
	isFinish := false

	pool.SetFinishCallback(func() {
		func(isFinish *bool) {
			*isFinish = true
		}(&isFinish)
	})

	pool.Start()
	for !isFinish {
		fmt.Print(2222)
	}
	result := pool.TaskResult()
	for i := range result {
		fmt.Print(result[i])
	}
	pool.Stop()
}

func download(url string) *res.Result {
	result := res.ResultInstance().SetResult(200, "", "")
	return result
}
