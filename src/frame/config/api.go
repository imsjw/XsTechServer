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
	if v == nil {
		return 0
	}

	switch v.(type) {
	case float64:
		return (int)(v.(float64))
	default:
		return v.(int)
	}
}

func GetInt64(key string) int64 {
	v := configInstance.Get(key)
	if v == nil {
		return 0
	}

	switch v.(type) {
	case float64:
		return (int64)(v.(float64))
	default:
		return v.(int64)
	}
}

func GetBool(key string) bool {
	v := configInstance.Get(key)
	if v == nil {
		return false
	}
	return v.(bool)
}
