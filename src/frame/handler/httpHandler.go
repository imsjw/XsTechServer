package handler

import (
	"common/src/protocol/http"
	"common/src/router"
	"io"
	"log"
	"net"
)

type HttpHandler struct {
}

func (This *HttpHandler) handler(conn net.Conn) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println("[http handler error] ", err)
			conn.Close()
		}
	}()
	//循环处理请求
	for {
		req, resp, err := http.Analysis(conn)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Println("[http protocol analysis error] ", err)
			return
		}

		i := router.GetInterface(req)
		if i == nil {
			log.Println("[url does not exist, url: ", req.GetURL(), "]")
			resp.SetStatusCode(404)
			resp.SetStatusMsg("Not Found")
			resp.SetHeader("Content-Type", "text/plain")
			resp.SetStrResult("404 URL Not Found")
		} else {
			i.Handler(req, resp, i)
		}

		http.OutResponse(req, resp, i)

		connection, exist := resp.GetHeader("Connection")
		if !exist || connection != "keep-alive" {
			conn.Close()
			return
		}
	}

}
