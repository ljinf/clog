package main

import (
	"github.com/ljinf/log"
	"log"
	"os"
)

func main() {
	clog.Info("std log")
	clog.SetOptions(clog.WithLevel(clog.DebugLevel))
	clog.Debug("change std log to debug level")

	//输出到文件
	file, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("create file test.log failed")
		return
	}
	defer file.Close()

	logger := clog.New(
		clog.WithLevel(clog.InfoLevel),
		clog.WithOutput(os.Stdout, file), //out to console and file
		clog.WithFormatter(&clog.JsonFormatter{IgnoreBasicFields: false}),
	)
	logger.Debug("debug log")
	logger.Info("custom log with json formatter")
	logger.Error("error log")
}
