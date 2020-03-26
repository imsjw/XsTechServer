package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type DefaultConfig map[string]interface{}

func (This *DefaultConfig) Get(key string) interface{} {
	return nil
}

func (This *DefaultConfig) Set(key string, value interface{}) {

}

func (This *DefaultConfig) Init() error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println("get config file path error [error info:", err, "]")
	}
	path := filepath.Join(dir, "..", DefaultConfigFilePath)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("read config file error [file path:", path, "] [error info:", err, "]")
		return err
	}

	err = json.Unmarshal(data, This)
	if err != nil {
		log.Println("config json unmarshal error [file path:", path, "] [error info:", err, "]")
		return err
	}
	log.Println("config init success [file path:", path, "]")
	return nil
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
