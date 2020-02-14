/*

 Copyright 2019 Advantech.
 Author: jin.xin@advaantech.com.cn.

*/
package main

import (
	"rest-shell/pkg/apis/restshell/v1alpha2"
	"rest-shell/pkg/db"
	"rest-shell/pkg/utils/syslog"
)

func main() {
	//stopCh := signals.SetupSignalHandler()

	//testing sso function
	//test.TestEiToken()
	//test.TestDevToken()

	//init db, for status save
	_, err := db.InitDB()
	if err != nil {
		LOG.Error("Fail to init db connection: %s")
	}

	// start rest api server
	v1alpha2.AddWebService()
	LOG.Fatal("restshellservice terminated")
}
