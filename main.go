package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/huoli2wan/crontab/commands"
	"github.com/huoli2wan/crontab/models/config"
)

func main() {
	var (
		err    error
		runCmd string
	)
	flag.StringVar(&runCmd, "name", "serve", "serve or worker")
	if runCmd == "serve" {
		if err = commands.Server(); err != nil {
			goto ERR
		}
		fmt.Println("master服务启动成功，端口号为：", config.G_config.ApiPort)
	} else if runCmd == "worker" {
		if err = commands.Worker(); err != nil {
			goto ERR
		}

		fmt.Println("worker服务启动成功，端口号为：", config.G_config.ApiPort)

	} else {
		err = errors.New("serve or worker")
		goto ERR
	}

	select {}
ERR:
	fmt.Println(err.Error())
}
