package main

import (
	"sync"
	"os"
	"flag"
	"option"
	"data"
	"collector"
	"web"
	"util"
)

func main() {
	//初始化日志
	Log := util.LoggerFactory()
	Log.Debug("解析配置文件")
	configFile := flag.String("c", "monitor.conf", "程序启动配置文件")
	flag.Parse()
	option, err := option.NewOption(*configFile)
	if nil != err {
		//todo logs
		Log.Debug("配置文件解析失败", err)
		os.Exit(1)
	}
	option.Log = Log
	stg := data.NewStorage()
	//启动url监控
	Log.Debug("启动URL健康检查")
	colt := collector.NewUrlCollector(stg, option)
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		defer w.Done()
		colt.Collect()

	}()
	//启动http服务
	httpServer := web.NewHttpServer(option)
	apiService := web.NewApiService(stg, option)
	apiService.Route(httpServer)
	
	w.Add(1)
	go func() {
		defer w.Done()
		httpServer.Start()
	}()
	w.Wait()
	//启动http服务
}
