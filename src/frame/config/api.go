package config

var configInstance Config

func init() {
	configInstance = new(DefaultConfig)
}

func SetInstance(instance Config) {
	configInstance = instance
}

func Init() error {
	return configInstance.Init()
}

func Get(key string) interface{} {
	return nil
}

func Set(key string, value interface{}) {

}
