package util

import (
	"log"
)

type LOG struct {
	
}

func LoggerFactory() *LOG {
	log.SetFlags(log.Ldate|log.Ltime)
	return &LOG{}
}

func(l *LOG)Debug(v ...interface{}) {
	log.Println(v)
}

func(l *LOG)Debugf(format string, v ...interface{}) {
	log.Printf(format, v...)
}