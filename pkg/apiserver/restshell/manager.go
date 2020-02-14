/*

 Copyright 2019 Advantech.
 Author: jin.xin@advaantech.com.cn.

*/

package restshell

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"rest-shell/pkg/models/manager"
	"rest-shell/pkg/utils/shell"
	"rest-shell/pkg/utils/syslog"
)

type ShellBody struct {
	Command string 		`json:"command,omitempty" description:"shell command"`
}

func GetStatus(req *restful.Request, resp *restful.Response) {
	var result out.OutputsResult
	result.Error = "ok"
	result.Status = http.StatusOK
	resp.WriteAsJson(&result)
	return
}

func RunBashCmd(request *restful.Request, response *restful.Response) {
	var body ShellBody
	var res *out.OutputsResult
	err := request.ReadEntity(&body)
	if err != nil {
		var result out.OutputsResult
		result.Status = http.StatusInternalServerError
		result.Error = "parse json body fail"
		response.WriteAsJson(result)
		return
	}
	command := body.Command
	if(command == ""){
		var result out.OutputsResult
		result.Status = http.StatusInternalServerError
		result.Error = "command string is null"
		response.WriteAsJson(result)
		return
	}
	LOG.Info("command: ", command)
	res = shell.PostBashCmd(command)
	response.WriteAsJson(res)
}

func RunShCmd(request *restful.Request, response *restful.Response) {
	var body ShellBody
	var res *out.OutputsResult
	err := request.ReadEntity(&body)
	if err != nil {
		var result out.OutputsResult
		result.Status = http.StatusInternalServerError
		result.Error = "parse json body fail"
		response.WriteAsJson(result)
		return
	}
	command := body.Command
	if(command == ""){
		var result out.OutputsResult
		result.Status = http.StatusInternalServerError
		result.Error = "command string is null"
		response.WriteAsJson(result)
		return
	}
	LOG.Info("command: ", command)
	res = shell.PostShCmd(command)
	response.WriteAsJson(res)
}

