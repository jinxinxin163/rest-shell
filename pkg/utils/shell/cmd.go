package shell

import (
	"bytes"
	"os/exec"
	"rest-shell/pkg/models/manager"
)


func PostBashCmd(command string) *out.OutputsResult {
	var result out.OutputsResult
	str, err := exec_bash(command)
    if err == nil {
		result.Status = 200
		result.Error = "none"
		result.Outputs = str
	} else {
		result.Status = 500
		result.Error = "error: run command failed"
		result.Outputs = ""
    }
	return &result
}

func exec_bash(s string) (string, error){
    cmd := exec.Command("/bin/bash", "-c", s)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    return out.String(), err
}


func PostShCmd(command string) *out.OutputsResult {
	var result out.OutputsResult
	str, err := exec_sh(command)
	if err == nil {
		result.Status = 200
		result.Error = "none"
		result.Outputs = str
	} else {
		result.Status = 500
		result.Error = "error: run command failed"
		result.Outputs = ""
	}
	return &result
}

func exec_sh(s string) (string, error){
	cmd := exec.Command("/bin/sh", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}
