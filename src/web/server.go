package web

import(
	"net/http"
	"option"
	"util"
)
type HttpServer struct {
	pubServer *http.Server
	pubMux *http.ServeMux
	priServer *http.Server
	priMux *http.ServeMux
	Log *util.LOG
}

func NewHttpServer(option *option.Option) *HttpServer{
	pubMux := http.NewServeMux()
	priMux := http.NewServeMux()
	
	pubServer := &http.Server{
		Addr:option.HttpPubServerAddr,
		Handler:pubMux,
	}
	
	priServer := &http.Server{
		Addr:option.HttpPriServerAddr,
		Handler:priMux,
	}
	
	return &HttpServer{
		pubServer:pubServer,
		pubMux:pubMux,
		priServer:priServer,
		priMux:priMux,
		Log:option.Log,
	}
	
}

func (httpServer *HttpServer)Start() {
	go httpServer.pubServer.ListenAndServe()
	httpServer.Log.Debugf("外网http服务启动，监听地址:[%s]",httpServer.pubServer.Addr)
	go httpServer.priServer.ListenAndServe()
	httpServer.Log.Debugf("内网http服务启动，监听地址:[%s]",httpServer.priServer.Addr)	
}
func (httpServer *HttpServer)PubRoute(
	route string, handler func(http.ResponseWriter, *http.Request)) {
	httpServer.pubMux.HandleFunc(route, handler)
}

func (httpServer *HttpServer)PriRoute(
	route string, handler func(http.ResponseWriter, *http.Request)) {
	httpServer.priMux.HandleFunc(route, handler)
}