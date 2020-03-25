package interfaces

type Interface struct {
	Handler InterfaceHandler
	url     string
	method  string
}

func (This *Interface) SetUrl(url string) {
	This.url = url
}

func (This *Interface) SetMethod(method string) {
	This.method = method
}

func NewInterface() *Interface {
	return new(Interface)
}
