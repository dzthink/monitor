package web

import(
    "data"
    "net/http"
)
type ApiServer struct {
    stg *data.Storage
    server *http.Server
}

func NewApiServer(option *option.Option, stg *data.Storage) *ApiServer {

}

func(apiServer *ApiServer)Serve() {

}

func(apiServer *ApiServer)route() map[string]func() {
    route := make(map[string]func())
    route["/health/url"] = func(w http.ResponseWriter, r *http.Request) {

    }
}

func(apiServer *ApiServer)urlHealth(w http.ResponseWriter, r *http.Request) {
    
}
