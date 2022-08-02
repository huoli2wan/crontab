package routes

import (
	"fmt"
	"github.com/huoli2wan/crontab/handers"
	"github.com/huoli2wan/crontab/models/config"
	"net"
	"net/http"
	"time"
)

//定义路由
func initRoute() *http.ServeMux {
	mux := http.NewServeMux()
	handler := handers.Handler{}
	mux.HandleFunc("/job/create", handler.CreateJob)
	mux.HandleFunc("/job/del", handler.DeleteJob)
	mux.HandleFunc("/job/list", handler.ListJob)
	mux.HandleFunc("/job/kill", handler.KillJob)
	return mux
}

//初始化服务
func InitApiServer() (err error) {
	var (
		mux        *http.ServeMux
		listener   net.Listener
		httpServer *http.Server
	)
	mux = initRoute()
	httpServer = &http.Server{
		Handler:      mux,
		ReadTimeout:  time.Duration(config.G_config.ApiReadTimeout) * time.Microsecond,
		WriteTimeout: time.Duration(config.G_config.ApiWriteTimeout) * time.Microsecond,
	}

	//启动TCP监听
	if listener, err = net.Listen("tcp", fmt.Sprintf(":%d", config.G_config.ApiPort)); err != nil {
		return
	}
	//启动了服务端
	go httpServer.Serve(listener)
	return
}
