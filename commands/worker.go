package commands

import (
	"github.com/huoli2wan/crontab/models/config"
	"github.com/huoli2wan/crontab/models/worker_manage"
)

func Worker() (err error) {

	//初始化线程
	initEnv()
	//加载配置
	err = config.InitConfig()
	// 服务注册
	err = worker_manage.InitRegister()
	return
}
