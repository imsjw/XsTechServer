package router

import (
	"frame/interfaces"
	"log"
)

var routerMap map[string]map[string]*interfaces.Interface

func init() {
	routerMap = make(map[string]map[string]*interfaces.Interface)
}

func Add(url string, handler interfaces.InterfaceHandler, method string) {
	if handler == nil {
		log.Println("router.Add error [handler is nil]")
		return
	}

	i := interfaces.NewInterface()
	i.Handler = handler
	i.SetUrl(url)
	i.SetMethod(method)

	_, exist := routerMap[method]
	if !exist {
		routerMap[method] = make(map[string]*interfaces.Interface)
	}

	_, exist = routerMap[method][url]
	if exist {
		log.Println("router.Add error [url exist] [method: ", method, " url: ", url, "]")
		return
	}

	routerMap[method][url] = i
	log.Println("router.Add method: ", method, " url: ", url)
}

func GetInterface(req interfaces.Request) *interfaces.Interface {

	urlMap, exist := routerMap[req.GetMethod()]
	if exist {
		return urlMap[req.GetURL()]
	}
	return nil
}
