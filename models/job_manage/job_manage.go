package job_manage

import (
	"context"
	"encoding/json"
	"github.com/huoli2wan/crontab/models/config"
	"github.com/huoli2wan/crontab/vars"
	"go.etcd.io/etcd/api/v3/mvccpb"

	clientv3 "go.etcd.io/etcd/client/v3"

	"time"
)

//任务管理器
type JobManage struct {
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

//任务
type Job struct {
	Name     string `json:"name"`      //任务名称
	Command  string `json:"command"`   //shell命令
	CronExpr string `json:"cron_expr"` //cron表达式
}

var (
	G_jobManage *JobManage
)

//初始化任务管理器
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
		DialTimeout: time.Duration(config.G_config.EtcdDialTimeout) * time.Microsecond, //连接超时
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

	defer client.Close()
	return
}

//Create 创建任务
func (jobManage *JobManage) Create(job *Job) (oldJob *Job, err error) {
	var (
		jobKey      string //etcd中保存的key
		jobValue    []byte //任务信息json
		putResponse *clientv3.PutResponse
		oldJobObj   Job
	)
	jobKey = vars.JOB_SAVE_DIR + job.Name
	if jobValue, err = json.Marshal(job); err != nil {
		return
	}
	//保存到etcd中
	if putResponse, err = jobManage.kv.Put(context.TODO(), jobKey, string(jobValue), clientv3.WithPrevKV()); err != nil {
		return
	}
	// 如果是更新，那么返回旧值
	if putResponse.PrevKv != nil {
		//对旧值做一个反序列化
		if err = json.Unmarshal(putResponse.PrevKv.Value, &oldJobObj); err != nil {
			err = nil
			return
		}
		oldJob = &oldJobObj

	}
	return
}

//列举任务
func (jobManage *JobManage) List() (jobList []*Job, err error) {
	var (
		dirKey  string
		getResp *clientv3.GetResponse
		job     *Job
		kvObj   *mvccpb.KeyValue
	)
	//任务保存的目录
	dirKey = vars.JOB_SAVE_DIR
	//获取任务下所有任务信息
	if getResp, err = jobManage.kv.Get(context.TODO(), dirKey, clientv3.WithPrefix()); err != nil {
		return
	}
	jobList = make([]*Job, 0)
	//遍历所有任务，进行反序列化
	for _, kvObj = range getResp.Kvs {
		job = &Job{}
		if err = json.Unmarshal(kvObj.Value, job); err != nil {
			continue
		}
		jobList = append(jobList, job)
	}

	return
}
