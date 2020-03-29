package handler

import (
	"frame/log"
	"frame/protocol/http"
	"frame/router"
	"net"
)

type HttpHandler struct {
}

func (This *HttpHandler) handler(conn net.Conn) {
	defer func() {
		err := recover()
		if err != nil {
			log.Error("frame.handler.HttpHandler 全局异常,错误信息:[", err, "]")
			conn.Close()
		}
	}()
	//循环处理请求
	for {
		req, resp := http.Analysis(conn)

		i := router.GetInterface(req)
		if i == nil {
			log.Error("[url does not exist, url: ", req.GetURL(), "]")
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
