package router

import (
	"fmt"
	"frame/interfaces"
	"frame/log"
)

var routerMap map[string]map[string]*interfaces.Interface

func init() {
	routerMap = make(map[string]map[string]*interfaces.Interface)
}

func Add(url string, handler interfaces.InterfaceHandler, method string) {
	if handler == nil {
		errMsg := fmt.Sprint("frame.router.Add 添加路由映射失败,处理处理函数为nil")
		log.Error(errMsg)
		panic(errMsg)
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
		errMsg := fmt.Sprint("frame.router.Add 添加路由映射失败,method和url已经存在,method:[", method, "] url:[", url, "]")
		log.Error(errMsg)
		panic(errMsg)
	}

	routerMap[method][url] = i
	log.Info("frame.router.Add 添加路由映射成功,method:[", method, "] url:[", url, "]")
}

func GetInterface(req interfaces.Request) *interfaces.Interface {

	urlMap, exist := routerMap[req.GetMethod()]
	if exist {
		return urlMap[req.GetURL()]
	}
	return nil
}
