package collector


import(
	"option"
	"data"
	"time"
	"fmt"
	"net/http"
)
type UrlCollector struct {
	urls []string
	timer *time.Ticker
	stg *data.Storage
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
			collector.detectUrl()
		}
	}
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
	fmt.Println(url)

	timebegin := time.Now().UnixNano()
	resp, err := http.Get(url)
	timeConsume := (time.Now().UnixNano() - timebegin)/1000000
	if err != nil {
		return false,timeConsume
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, timeConsume
	}
	fmt.Println(url + ":" + fmt.Sprintf("%d", timeConsume))
	return true, timeConsume
}
