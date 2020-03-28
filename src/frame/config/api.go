package config

var configInstance Config

func init() {
	configInstance = NewDefaultConfig()
}

func SetInstance(instance Config) {
	configInstance = instance
}

func Init() {
	configInstance.Init()
}

func Get(key string) interface{} {
	return nil
}

func Set(key string, value interface{}) {

}
