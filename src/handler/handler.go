package handler

import "net"

type Handler interface {
	handler(conn net.Conn)
}
