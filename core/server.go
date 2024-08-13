package core

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"go.uber.org/zap"
	"gva/global"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.Gxva_CONFIG.System.UseMultipoint || global.Gxva_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
	if global.Gxva_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// 从db加载jwt数据
	if global.Gxva_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.Gxva_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.Gxva_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 gin-xjjx-vue-admin
	当前版本:v2.6.5
    加群方式:微信号：gsz19831210
	项目地址：
	插件市场:
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	--------------------------------------版权声明--------------------------------------
`, address)
	global.Gxva_LOG.Error(s.ListenAndServe().Error())
}
