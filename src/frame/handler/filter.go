package handler

import "frame/interfaces"

//返回false 代表拦截,返回true代表不拦截
type Filter func(interfaces.Request, interfaces.Response, *interfaces.Interface) bool

func isFilter(req interfaces.Request, resp interfaces.Response, i *interfaces.Interface) bool {
	//判断是否是白名单中的url
	urls, exist := filterWhiteList[req.GetMethod()]
	if exist {
		for _, url := range urls {
			if url == req.GetURL() {
				return true
			}
		}
	}

	//调用过滤器函数判断是否过滤
	for _, filter := range filters {
		if !filter(req, resp, i) {
			return false
		}
	}
	return true
}
