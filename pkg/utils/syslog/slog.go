/*

 Copyright 2019 Advantech.
 Author: jin.xin@advaantech.com.cn.

*/
package LOG

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func init(){
	//log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

func Panic(args ...interface{}){
	log.Panic(args)
}
func Fatal(args ...interface{}){
	log.Fatal(args)
}
func Error(args ...interface{}){
	log.Error(args)
}
func Warn(args ...interface{}){
	log.Warn(args)
}
func Info(args ...interface{}){
	log.Info(args)
}
func Debug(args ...interface{}){
	log.Debug(args)
}
