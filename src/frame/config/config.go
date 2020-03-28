package config

type Config interface {
	Init()
	Set(key string, value interface{})
}
