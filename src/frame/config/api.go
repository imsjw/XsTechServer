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

func Exist(key string) bool {
	return configInstance.Exist(key)
}

func Set(key string, value interface{}) {
	configInstance.Set(key, value)
}

func Get(key string) interface{} {
	return configInstance.Get(key)
}

func GetString(key string) string {
	v := configInstance.Get(key)
	if v != nil {
		return v.(string)
	}
	return ""
}

func GetInt(key string) int {
	v := configInstance.Get(key)
	if v != nil {
		return (int)(v.(float64))
	}
	return 0
}
