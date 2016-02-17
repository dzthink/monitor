package data

type Storage struct {
    urlStatus map[string]int32
}

func(stg *Storage)RefreshUrlStatus(url string, status int32) {
    stg.urlStatus[url] = status
}
