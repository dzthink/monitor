package main

import (
	"sync"
	"os"
	"flag"
	"fmt"
	"option"
	"data"
	"collector"
)

func main() {
	configFile := flag.String("c", "monitor.conf", "程序启动配置文件")
	fmt.Println("monitor starting")
	option, err := option.NewOption(*configFile)
	if nil != err {
		//todo logs
		os.Exit(1)
	}
	
	stg := data.NewStorage()
	//启动url监控
	colt := collector.NewUrlCollector(stg, option)
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		defer w.Done()
		colt.Collect()

	}()
	w.Wait()
	//启动http服务
}