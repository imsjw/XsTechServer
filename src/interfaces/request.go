package interfaces

type Request interface {
	GetURL() string
	GetMethod() string
	GetObjParam(v interface{}) error
}
