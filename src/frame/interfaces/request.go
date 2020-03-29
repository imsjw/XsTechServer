package interfaces

type Request interface {
	GetURL() string
	GetMethod() string
	GetObjParam(v interface{}) error
	GetHeader(k string) (string, bool)
}
