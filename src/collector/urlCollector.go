package collector


import(
	"option"
	"data"
	"time"
	"net/http"
	"util"
)
type UrlCollector struct {
	urls []string
	timer *time.Ticker
	stg *data.Storage
	Log *util.LOG
}

func NewUrlCollector(stg *data.Storage, option *option.Option) *UrlCollector{
	urlCollector := &UrlCollector{
		timer:time.NewTicker(time.Duration(option.HttpTimeInterval)*time.Second),
		urls:option.HttpUrlList,
		stg:stg,
	}
	return urlCollector
}
func(collector *UrlCollector) Collect() {
	for {
		select {
		case <-collector.timer.C:
			go collector.detectUrl()
		}
	}
	collector.Log.Debug("url健康检查已启动")
}

func(collector *UrlCollector)detectUrl() {
	for i := 0; i < len(collector.urls); i++ {
		url := collector.urls[i]
		health, latency := collector.getUrlStatus(url)
		urlHealth := data.UrlHealth{Health:health,Latency:latency}
		collector.stg.RefreshUrlStatus(url, urlHealth)
	}
}

func (collector *UrlCollector)getUrlStatus(url string) (bool,int64){
	timebegin := time.Now().UnixNano()
	resp, err := http.Get(url)
	timeConsume := (time.Now().UnixNano() - timebegin)/1000000
	if err != nil {
		collector.Log.Debugf("检查url[%s]时发生异常,错误信息[%s]", url, err)
		return false,timeConsume
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		collector.Log.Debugf("url[%s]状态异常,状态码[%d]", url, resp.StatusCode)
		return false, timeConsume
	}
	collector.Log.Debugf("url[%s]状态正常，时延[%d]", url, timeConsume)
	return true, timeConsume
}
