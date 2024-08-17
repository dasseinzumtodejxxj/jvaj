package main

import (
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"gva/core"
	"gva/global"
	"gva/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.Gxva_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.Gxva_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.Gxva_LOG)
	global.Gxva_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.Gxva_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.Gxva_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
