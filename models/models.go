package models

import "encoding/json"

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
