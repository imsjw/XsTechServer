package app

import (
	"fmt"
	"frame/config"
	"frame/handler"
	"frame/log"
	"net"
)

func listen() {
	if !config.Exist(ConfigKeyHttpServerPort) {
		errMsg := fmt.Sprint("frame.app.listen 缺少配置:[", ConfigKeyHttpServerPort, "]")
		log.Error(errMsg)
		panic(errMsg)
	}
	port := config.GetInt(ConfigKeyHttpServerPort)
	listener, err := net.Listen("tcp", fmt.Sprint("localhost:", port))
	if err != nil {
		log.Error("frame.app.listen 端口监听失败,端口:[", port, "],错误信息:[", err, "]")
		panic(err)
	}

	log.Info("frame.app.listen 端口监听成功,端口:[", port, "]")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handler.Http(conn)
	}
}
