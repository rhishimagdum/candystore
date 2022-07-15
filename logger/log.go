package logger

import (
	"sync"

	"github.com/rhishimagdum/candystore/config"
	"github.com/sirupsen/logrus"
)

var Logger logrus.Logger = *logrus.New()

var once sync.Once

func Log() *logrus.Logger {
	once.Do(func() {
		lvl, _ := logrus.ParseLevel(config.GetConfig().LogLevel)
		Logger.SetLevel(lvl)
	})
	return &Logger
}
