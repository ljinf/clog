package main

import (
	"github.com/ljinf/log"
	"os"
)

func main() {
	log.Info("std log")
	log.SetOptions(log.WithLevel(log.DebugLevel))
	log.Debug("change std log to debug level")

	//输出到文件
	file, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("create file test.log failed")
		return
	}
	defer file.Close()

	logger := log.New(
		log.WithLevel(log.InfoLevel),
		log.WithOutput(os.Stdout, file), //out to console and file
		log.WithFormatter(&log.JsonFormatter{IgnoreBasicFields: false}),
	)
	logger.Debug("debug log")
	logger.Info("custom log with json formatter")
	logger.Error("error log")
}
