package main

import (
	"backend-pack/internal/app/core"
	"backend-pack/internal/app/global"
	"go.uber.org/zap"
)

func main() {
	// 初始化Viper
	core.Viper()
	// 初始化zap日志库
	global.LOG = core.Zap()
	// 注册全局logger
	zap.ReplaceGlobals(global.LOG)
	core.RunWindowsServer()
}
