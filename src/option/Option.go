package option

import (
	"strconv"
	"errors"
	"reflect"
	"strings"
	"fmt"
	"io"
	"bufio"
	"os"
	"util"
)

type Option struct{
	HttpTimeInterval int "httpTimeInterval"
	HttpUrlList []string "httpUrl"
	HttpPubServerAddr string "httpPubServerAddr"
	HttpPriServerAddr string "httpPriServerAddr"
	Log *util.LOG	
}

func NewOption(path string) (*Option, error) {
	//默认配置
	option := Option{
		HttpTimeInterval:5,
	}
	
	//读取文件
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

func (op *Option) SetHttpTimeInterval(value string) error{
	interval, err := strconv.Atoi(value)
	if err != nil {
		//todo 参数值错误
		return err
	}
	op.HttpTimeInterval = interval
	return nil
}

func (op *Option) SetHttpUrl(value string) error{
	op.HttpUrlList = append(op.HttpUrlList, value)
	return nil
}

func(op *Option) parseLine(line string) {
	//todo 过滤空格过滤注释行
	confNameValue := strings.Split(line, "=")
	if len(confNameValue) != 2 {
		//todo logs
		fmt.Println("配置格式错误!")
		return;
	}
	confName := strings.TrimSpace(confNameValue[0])
	confVal := strings.TrimSpace(confNameValue[1])
	setter := "Set" + strings.Title(confName)

	opValue := reflect.ValueOf(op)
	if opSetter := opValue.MethodByName(setter);opSetter.IsValid() {//找到setter方法
		param := make([]reflect.Value, 1)
		param[0] = reflect.ValueOf(confVal)
		opSetter.Call(param)
	} else if fieldIndex, err := op.searchOption(confName); err == nil {//找到同名tag属性
		//未处理的panic
		opValue.Elem().Field(fieldIndex).SetString(confVal)
	} else {//不支持的配置
		//todo logs	
		fmt.Println("配置不支持")
	}
}

func(op *Option)searchOption(name string)(int, error){
	opValue := reflect.TypeOf(op).Elem()
	fieldNum := opValue.NumField()
	for i := 0; i < fieldNum; i++ {
		if strings.EqualFold(string(opValue.Field(i).Tag), name) {
			return i, nil
		}
	}
	return -1, errors.New("field for config:" + name + " not found!")
}
