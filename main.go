package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"workspace/ginweb/pkg/mongodb"
	"workspace/ginweb/pkg/util"

	//"os"
	//"os/signal"
	//"syscall"

	//"syscall"
	"workspace/ginweb/conf"
	"workspace/ginweb/pkg/logs"
	//"workspace/ginweb/pkg/util"
	"workspace/ginweb/router"
)

func main() {
	//initialize log
	logs.Init("./log", "log", 3, false)
	//initialize mongo
	err := mongodb.Init()
	if err != nil {
		logs.Errorf("mongo init err:%+v", err)
	}

	// catch exit signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	// Start service monitoring
	go func(port int, mode string) {
		if err := router.RouterRun(port, mode); err != nil {
			log.Printf("server run failed:%v", err)
			return
		}
	}(conf.HTTP_PORT, "dev")

	select {}
	log.Printf("server[%s] serving on %s:%d", "dev", util.GetLocalIPv4(), conf.HTTP_PORT)
	s := <-sig
	log.Printf("receive signal:%d , server stopping ...", s)
}
