package main

import (
	"github.com/aphly/go_utils"
	"time"
)

func main() {
	NewLog := go_utils.NewLog("log", "mysql", 1)
	NewLog.Info("xxx")
	NewLog.Info(12)
	NewLog.Info(time.Now().Unix())
}
