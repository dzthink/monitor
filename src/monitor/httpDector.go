package monitor

import(
	"option"
)
type UrlDector struct {
	interval int
	urls map[string][]string
	
}

func NewUrlDector(option *option.Option) *UrlDector{
	UrlDector := &UrlDector{
		interval:option.HttpTimeInterval,
	}
	
}
func(dector *UrlDector) start() {
	
}

