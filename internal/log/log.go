package log

import (
	"log"
	"os"
	"sync"
)

type logger struct {
	logger *log.Logger
}

var instance *logger

var once sync.Once

func GetInstance() *logger {
	once.Do(func() {
		file, err := os.Create("./log.log")
		if err != nil {
			panic(err)
		}
		instance = &logger{
			logger: log.New(file, "", log.LstdFlags),
		}
	})
	return instance
}

func (l *logger) Info(msg string) {
	l.logger.Println("INFO: " + msg)
}

func (l *logger) Error(format string, err error) {
	l.logger.Fatalf(format, err)
}
