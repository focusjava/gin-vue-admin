package core

import (
	"fmt"
	"time"

	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"gin-vue-admin/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
欢迎使用 gin-vue-admin
当前版本:v2.5.7
插件市场:https://plugin.gin-vue-admin.com
GVA讨论社区:https://support.qq.com/products/371961
默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
默认前端文件运行地址:http://127.0.0.1:8080
项目介绍：https://www.gin-vue-admin.com/guide/introduce/project.html
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
