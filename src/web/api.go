package web

import (
	"fmt"
)

import(
    "data"
    "net/http"
	"util"
	"option"
)
type ApiService struct {
    stg *data.Storage
	Log *util.LOG
}

func NewApiService(stg *data.Storage, option *option.Option) *ApiService {
	return &ApiService{
		stg:stg,
		Log:option.Log,
	}
}


func(apiService *ApiService)Route(httpServer *HttpServer) {
	//输出url健康状态
    httpServer.PriRoute("/health/url", func(w http.ResponseWriter, r *http.Request) {
		apiService.urlHealth(w, r)
    })
}


func(apiService *ApiService)urlHealth(w http.ResponseWriter, r *http.Request) {
   	urlStatus := apiService.stg.GetUrlStatus()
	output := ""	
	for key, value := range urlStatus {
		okStr := "fail"
		if value.Health {
			okStr = "ok"
		}
		output += key + " " + okStr + " " + fmt.Sprintf("%d", value.Latency) + "\n"
	}
	w.Write([]byte(output))
}
