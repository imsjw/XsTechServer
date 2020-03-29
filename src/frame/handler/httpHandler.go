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
			log.Warning("[frame.handler.HttpHandler 路由中不存在映射,method:[", req.GetMethod(), "] url:[", req.GetURL(), "]")
			resp.SetStatusCode(404)
			resp.SetStatusMsg("Not Found")
			resp.SetHeader("Content-Type", "text/plain")
			resp.SetStrResult("404 URL Not Found")
		} else {
			if isFilter(req, resp, i) {
				i.Handler(req, resp, i)
			} else {
				log.Warning("[frame.handler.HttpHandler 请求被拦截,method:[", req.GetMethod(), "] url:[", req.GetURL(), "]")
				resp.SetStatusCode(403)
				resp.SetStatusMsg("Permission denied")
				resp.SetHeader("Content-Type", "text/plain")
				resp.SetStrResult("403 Permission denied")
			}
		}

		http.OutResponse(req, resp, i)

		connection, exist := resp.GetHeader("Connection")
		if !exist || connection != "keep-alive" {
			conn.Close()
			return
		}
	}
}
