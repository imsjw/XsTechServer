package config

import (
	"encoding/json"
	"fmt"
	"frame/log"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DefaultConfig struct {
	kv map[string]interface{}
}

func NewDefaultConfig() *DefaultConfig {
	res := new(DefaultConfig)
	res.kv = make(map[string]interface{})
	return res
}

func (This *DefaultConfig) Init() {
	//获取文件路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Error("frame.config.DefaultConfig,Init 获取配置文件路径错误,错误信息:[", err, "]")
		panic(err)
	}
	path := filepath.Join(dir, "..", DefaultConfigFilePath)

	//读取配置文件
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("frame.config.DefaultConfig.Init 读取配置文件错误,文件路径:[", path, "],错误信息:[", err, "]")
		panic(err)
	}

	var jsonMap map[string]interface{}

	//配置文件转json
	err = json.Unmarshal(data, &jsonMap)
	if err != nil {
		log.Error("frame.config.DefaultConfig.Init JSON解析失败 文件路径:[", path, "],错误信息:[", err, "]")
		panic(err)
	}

	This.jsonToKeyValue(jsonMap, "")

	log.Info("frame.config.DefaultConfig.Init 配置初始化成功,文件路径:[", path, "]")
}

func (This *DefaultConfig) jsonToKeyValue(jsonMap map[string]interface{}, prefix string) {
	for k, v := range jsonMap {
		key := fmt.Sprint(prefix, ".", k)
		if prefix == "" {
			key = key[1:len(key)]
		}

		switch v.(type) {
		case map[string]interface{}:
			This.jsonToKeyValue((v.(map[string]interface{})), key)
		default:
			This.kv[key] = v
		}
	}
}

func (This *DefaultConfig) Exist(key string) bool {
	_, exist := This.kv[key]
	return exist
}

func (This *DefaultConfig) Set(key string, value interface{}) {
	This.kv[key] = value
}

func (This *DefaultConfig) Get(key string) interface{} {
	return This.kv[key]
}
