package inits

import (
	"github.com/yungsem/gotool/log"
	"os"
)

var Log *log.Log

// initLog 初始化 Log
func initLog() {
	Log = log.NewLog(Conf.Log.Output, Conf.Log.Level, Conf.Log.Path + string(os.PathSeparator) + Conf.Server.Name )
}
