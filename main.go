package main

import (
	"fmt"
	"github.com/fanyiheng/go-web-demo/logging"
	"github.com/fanyiheng/go-web-demo/persist"
	"github.com/fanyiheng/go-web-demo/router"
	"github.com/fanyiheng/go-web-demo/setting"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	//todo 使用flag读取配置文件路径
	setting.Setup("")
	persist.Setup()
	logging.Setup()
	routersInit := router.Setup()
	readTimeout := setting.ServerSetting.ReadTimeout * time.Second
	writeTimeout := setting.ServerSetting.WriteTimeout * time.Second
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	logrus.Infof("start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
