package app

import (
	"common/src/config"
	"common/src/data/source"
	"common/src/handler"
	"fmt"
	"log"
	"net"
)

func Start() error {
	config.Init()

	source.Init()

	listener, err := net.Listen("tcp", fmt.Sprint("localhost:", config.GetHttpServerPort()))
	if err != nil {
		log.Println("net listener error [error info:", err, "]")
		return err
	}
	log.Println("http listener success [port:", config.GetHttpServerPort(), "]")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handler.Http(conn)
	}

}
