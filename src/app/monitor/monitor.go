package main

import (
	"os"
	"flag"
	"fmt"
	"option"
)

func main() {
	configFile := flag.String("c", "monitor.conf", "程序启动配置文件")
	fmt.Println("monitor starting")
	option, err := option.NewOption(*configFile)
	if nil != err {
		//todo logs
		os.Exit(1)
	}
	
	//启动http监控
}