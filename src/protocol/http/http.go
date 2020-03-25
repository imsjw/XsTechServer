package http

import (
	"net"
)

type Http struct {
	conn            net.Conn
	method          string
	url             string
	protocolVersion string
	header          map[string]string
	body            []byte
}

func (This *Http) GetConn() net.Conn {
	return This.conn
}

func (This *Http) GetMethod() string {
	return This.method
}

func (This *Http) GetURL() string {
	return This.url
}

func (This *Http) GetProtocolVersion() string {
	return This.protocolVersion
}

func (This *Http) GetHeader(k string) (string, bool) {
	v, exist := This.header[k]
	return v, exist
}

func (This *Http) SetHeader(k string, v string) {
	This.header[k] = v
}

func (This *Response) GetBody() []byte {
	return This.body
}

func (This *Response) SetBody(body []byte) {
	This.body = body
}
