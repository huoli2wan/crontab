package handers

import (
	"encoding/json"
	"github.com/huoli2wan/crontab/models"
	"github.com/huoli2wan/crontab/models/job_manage"
	"net/http"
)

type Handler struct {
}

func (*Handler) ListJob(res http.ResponseWriter, req *http.Request) {
	var (
		jobList []*job_manage.Job
		err     error
		bytes   []byte
	)
	if jobList, err = job_manage.G_jobManage.List(); err != nil {
		goto ERR
	}

	if bytes, err = models.BuildResponse(0, "success", jobList); err == nil {
		res.Write(bytes)
		return
	}

ERR:
	if bytes, err = models.BuildResponse(-1, err.Error(), nil); err == nil {
		res.Write(bytes)
	}

}

func (*Handler) CreateJob(res http.ResponseWriter, req *http.Request) {
	var (
		err     error
		bytes   []byte
		postJob string
		job     job_manage.Job
		oldJob  *job_manage.Job
	)
	if err = req.ParseForm(); err != nil {
		goto ERR
	}
	postJob = req.PostForm.Get("job")
	if err = json.Unmarshal([]byte(postJob), &job); err != nil {
		goto ERR
	}
	//保存到etcd中
	if oldJob, err = job_manage.G_jobManage.Create(&job); err != nil {
		goto ERR
	}
	//返回正常应答
	if bytes, err = models.BuildResponse(0, "success", oldJob); err == nil {
		res.Write(bytes)
		return
	}

ERR:
	//返回异常应答
	if bytes, err = models.BuildResponse(-1, err.Error(), nil); err == nil {
		res.Write(bytes)
	}
}

func (*Handler) DeleteJob(res http.ResponseWriter, req *http.Request) {
	var (
		err     error
		oldJob  *job_manage.Job
		bytes   []byte
		jobName string
	)
	if err = req.ParseForm(); err != nil {
		goto ERR
	}
	jobName = req.PostForm.Get("name")

	if oldJob, err = job_manage.G_jobManage.DeleteJob(jobName); err != nil {
		goto ERR
	}
	if bytes, err = models.BuildResponse(0, "success", oldJob); err == nil {
		res.Write(bytes)
		return
	}

ERR:
	if bytes, err = models.BuildResponse(-1, err.Error(), nil); err == nil {
		res.Write(bytes)
	}
}

func (*Handler) KillJob(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("KillJob"))
	return
}
