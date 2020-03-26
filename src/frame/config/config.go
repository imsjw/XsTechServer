package config

type Config interface {
	Init() error
	Get(key string) interface{}
	Set(key string, value interface{})
}
