package main

import (
	"fmt"
	"github.com/huoli2wan/crontab/models/config"
)

func main() {
	var (
		err error
	)
	if err = config.InitConfig(); err != nil {
		panic(err)
	}
	fmt.Println(config.G_config)
}
