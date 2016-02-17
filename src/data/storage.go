package data

type UrlHealth struct {
	Health bool
	Latency int64
}

type Storage struct {
    urlStatus map[string]UrlHealth
}

func NewStorage() *Storage{
	stg := &Storage{
		urlStatus : make(map[string]UrlHealth),
	}
	return stg
}
func(stg *Storage)RefreshUrlStatus(url string, urlHealth UrlHealth) {
    stg.urlStatus[url] = urlHealth
}
