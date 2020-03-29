package handler

import (
	"net"
)

var httpHandlerInstance Handler = new(HttpHandler)
var filters []Filter

func Http(conn net.Conn) {
	httpHandlerInstance.handler(conn)
}

func AddFilter(filter Filter) {
	filters = append(filters, filter)
}
