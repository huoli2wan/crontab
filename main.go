package main

import (
	"fmt"
	"github.com/huoli2wan/crontab/models/config"
	"github.com/huoli2wan/crontab/models/job_manage"
	"github.com/huoli2wan/crontab/routes"
	"runtime"
	"time"
)

// initEnv 初始化线程数量
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	var (
		err error
	)
	//初始化线程
	initEnv()
	//加载配置
	if err = config.InitConfig(); err != nil {
		goto ERR
	}
	//初始化任务管理器
	if err = job_manage.InitJobManage(); err != nil {
		goto ERR
	}
	//启动API,HTTP服务
	if err = routes.InitApiServer(); err != nil {
		goto ERR
	}

	// 正常退出
	for {
		time.Sleep(1 * time.Second)
	}
	return

ERR:
	fmt.Println(err)
}
