package handler

import (
	"net"
)

var httpHandlerInstance Handler = new(HttpHandler)
var filters []Filter
var filterWhiteList map[string][]string

func init() {
	filterWhiteList = make(map[string][]string)
}

func Http(conn net.Conn) {
	httpHandlerInstance.handler(conn)
}

func AddFilter(filter Filter) {
	filters = append(filters, filter)
}

func AddFilterWhiteList(method string, url string) {
	urls, exist := filterWhiteList[method]
	if !exist {
		urls = []string{}
	}
	urls = append(urls, url)
	filterWhiteList[method] = urls
}
