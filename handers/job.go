package handers

import "net/http"

type Handler struct {
}

func (*Handler) ListJob(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("ListJob"))
	return
}

func (*Handler) CreateJob(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("CreateJob"))
	return
}

func (*Handler) DeleteJob(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("DeleteJob"))
	return
}

func (*Handler) KillJob(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("KillJob"))
	return
}
