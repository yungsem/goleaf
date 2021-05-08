package main

import (
	"fmt"
	inits2 "github.com/yungsem/goleaf/inits"
	router2 "github.com/yungsem/goleaf/router"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic value: %v\n", r)
		}
	}()

	// 初始化路由
	r := router2.Init()

	go func() {
		inits2.Log.Debug("服务启动，开启 pprof 采样")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()


	// 运行服务
	inits2.Log.Debug("服务启动，监听端口%s", inits2.Conf.Server.Port)
	err := r.Run(":" + inits2.Conf.Server.Port)
	if err != nil {
		panic(err)
	}
}
