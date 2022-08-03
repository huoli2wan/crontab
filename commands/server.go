package commands

import (
	"github.com/huoli2wan/crontab/models/config"
	"github.com/huoli2wan/crontab/models/job_manage"
	"github.com/huoli2wan/crontab/routes"
	"runtime"
)

// initEnv 初始化线程数量
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func Server() (err error) {

	//初始化线程
	initEnv()
	//加载配置
	err = config.InitConfig()
	//初始化任务管理器
	err = job_manage.InitJobManage()
	//启动API,HTTP服务
	err = routes.InitApiServer()
	return
}
