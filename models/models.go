package models

import (
	"encoding/json"
	"github.com/huoli2wan/crontab/models/config"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	EtcdClient *EtcdConf
)

//etcd客户端初始化
type EtcdConf struct {
	Client *clientv3.Client
	Kv     clientv3.KV
	Lease  clientv3.Lease
}

// HTTP接口应答
type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Retcode int         `json:"retcode"`
}

//应答方法
func BuildResponse(retcode int, message string, data interface{}) (resp []byte, err error) {
	var (
		response Response
	)
	response.Data = data
	response.Message = message
	response.Retcode = retcode
	//序列换json
	resp, err = json.Marshal(response)
	return
}

//初始化etcd客户端
func InitJobManage() (err error) {
	var (
		conf   clientv3.Config
		client *clientv3.Client
		kv     clientv3.KV
		lease  clientv3.Lease
	)
	//初始化配置
	conf = clientv3.Config{
		Endpoints:   config.G_config.EtcdEndpoints,                                     //集群地址
		DialTimeout: time.Duration(config.G_config.EtcdDialTimeout) * time.Millisecond, //连接超时
	}
	//建立连接
	if client, err = clientv3.New(conf); err != nil {
		return
	}
	// 得到KV和Lease的API子集
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)
	//赋值单例
	G_jobManage = &JobManage{
		client: client,
		kv:     kv,
		lease:  lease,
	}

	//defer client.Close()
	return
}
