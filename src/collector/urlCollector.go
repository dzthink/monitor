package monitor

import(
	"option"
	"data"
	"time"
	"http"
)
type UrlCollector struct {
	urls []string
	timer *time.Ticker
}

func NewUrlCollector(option *option.Option) *UrlCollector{
	urlCollector := &UrlCollector{
		timer:time.NewTicker(option.HttpTimeInterval*time.Second),
		urls:option.HttpUrlList

	}
	return &urlCollector
}
func(collector *UrlCollector) Collect(stg *data.Storage) {
	for {
		select {
		case <-collector.timer.C:
			go collector.detectUrl(stg)
		}
	}
}

func(collector *UrlCollector)detectUrl(stg *data.Storage) {
	for i := 0; i < len(collector.urls); i++ {

	}
}

func (collector *UrlCollector)getUrlStatus(url string) (bool,int32){
	timebegin := time.Now().Unix()
	resp, err := http.Get(url)
	timeConsume := time.Now().Unix() - timebegin
	if err != nil {
		return false,timeConsume
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, timeConsume
	}

	return true, timeConsume
}
