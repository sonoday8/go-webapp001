package config

import (
	"os"

	"github.com/labstack/echo/v4/middleware"
	"github.com/sonoday8/webapp001/app/env"
)

// AccessLogConfig return middleware.LoggerConfig
var AccessLogConfig = middleware.LoggerConfig{
	Output: getLogFile(env.GetStr("ACCESS_LOG", "logs/access.log")),
}

func getLogFile(file string) *os.File {
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic("log file error")
	}
	return logFile
}
