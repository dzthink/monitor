package option

import (
	"strings"
	"fmt"
	"io"
	"bufio"
	"os"
)

type Option struct{
	HttpTimeInterval int32 "httpTimeInterval"
	HttpUrlList []string "http.url"
}

func NewOption(path string) (*Option, error) {
	//默认配置
	option := Option{
		HttpTimeInterval:5,
	}
	
	//读取文件
	fmt.Println(path)
	fhandle, err := os.Open(path)
	if nil != err {
		return &option, err
	}
	reader := bufio.NewReader(fhandle)
	for {
		line, err := reader.ReadString('\n')
		option.parseLine(strings.Trim(line,"\n"))
		if err != nil || io.EOF == err {
			//文件解析完毕
			break
		}
	}
	return &option, nil
}


func(op *Option) parseLine(line string) {
	
}