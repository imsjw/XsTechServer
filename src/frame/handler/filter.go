package handler

import "frame/interfaces"

//返回false 代表拦截,返回true代表不拦截
type Filter func(interfaces.Request, interfaces.Response, *interfaces.Interface) bool

func isFilter(req interfaces.Request, resp interfaces.Response, i *interfaces.Interface) bool {
	for _, filter := range filters {
		if !filter(req, resp, i) {
			return false
		}
	}
	return true
}
