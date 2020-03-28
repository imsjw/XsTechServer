package config

type Config interface {
	Init()
	Exist(key string) bool
	Set(key string, value interface{})
	Get(key string) interface{}
}
