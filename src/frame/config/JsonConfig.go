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
		log.Error("frame.config.DefaultConfig,Init 读取配置文件错误,文件路径:[", path, "],错误信息:[", err, "]")
		panic(err)
	}

	var jsonMap map[string]interface{}

	//配置文件转json
	err = json.Unmarshal(data, &jsonMap)
	if err != nil {
		log.Error("frame.config.DefaultConfig,Init JSON解析失败 文件路径:[", path, "],错误信息:[", err, "]")
		panic(err)
	}

	This.jsonToKeyValue(jsonMap, "")

	log.Info("frame.config.DefaultConfig,Init 配置初始化成功,文件路径:[", path, "]")
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

func (This *DefaultConfig) Get(key string) interface{} {
	return nil
}

func (This *DefaultConfig) Set(key string, value interface{}) {

}

//func (This *DefaultConfig) GetHttpServerPort() int {
//	return This.Http.Server.Port
//}
//
//func (This *DefaultConfig) GetHttpOAuthEnable() bool {
//	return This.Http.OAuth.Enable
//}
//
//func (This *DefaultConfig) GetHttpOAuthPasswordSalt() string {
//	return This.Http.OAuth.Password.Salt
//}
//
//func (This *DefaultConfig) GetDataSourceDriverName() string {
//	return This.DataSource.DriverName
//}
//
//func (This *DefaultConfig) GetDataSourceUserName() string {
//	return This.DataSource.UserName
//}
//func (This *DefaultConfig) GetDataSourcePassword() string {
//	return This.DataSource.Password
//}
//func (This *DefaultConfig) GetDataSourceHost() string {
//	return This.DataSource.Host
//}
//func (This *DefaultConfig) GetDataSourceDBName() string {
//	return This.DataSource.DBName
//}
//
//func (This *DefaultConfig) GetDataSourcePort() int {
//	return This.DataSource.Port
//}
//
//func (This *DefaultConfig) GetHttpOAuthAccessTokenSalt() string {
//	return This.Http.OAuth.AccessToken.Salt
//}
//
//func (This *DefaultConfig) GetHttpOAuthAccessTokenValidTime() int64 {
//	return This.Http.OAuth.AccessToken.ValidTime
//}
//
//func (This *DefaultConfig) GetHttpOAuthRefreshTokenSalt() string {
//	return This.Http.OAuth.RefreshToken.Salt
//}
//
//func (This *DefaultConfig) GetHttpOAuthRefreshTokenValidTime() int64 {
//	return This.Http.OAuth.RefreshToken.ValidTime
//}
