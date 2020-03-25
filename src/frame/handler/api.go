package handler

import "net"

var httpHandlerInstance Handler = new(HttpHandler)

func Http(conn net.Conn) {
	httpHandlerInstance.handler(conn)
}
