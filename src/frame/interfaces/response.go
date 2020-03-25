package interfaces

type Response interface {
	GetURL() string
	SetObjResult(interface{}) error
	SetStrResult(string)
}
