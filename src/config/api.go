package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var configInstance Config

func Init() error {
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

	defaultConfig := new(DefaultConfig)
	err = json.Unmarshal(data, defaultConfig)
	if err != nil {
		log.Println("config json unmarshal error [file path:", path, "] [error info:", err, "]")
		return err
	}
	configInstance = defaultConfig
	log.Println("config init success [file path:", path, "]")
	return nil
}

func GetHttpServerPort() int {
	return configInstance.GetHttpServerPort()
}

func GetOAuthEnable() bool {
	return configInstance.GetHttpOAuthEnable()
}

func GetHttpOAuthPasswordSalt() string {
	return configInstance.GetHttpOAuthPasswordSalt()
}

func GetDataSourceDriverName() string {
	return configInstance.GetDataSourceDriverName()
}

func GetDataSourceUserName() string {
	return configInstance.GetDataSourceUserName()
}

func GetDataSourcePassword() string {
	return configInstance.GetDataSourcePassword()
}

func GetDataSourceHost() string {
	return configInstance.GetDataSourceHost()
}

func GetDataSourcePort() int {
	return configInstance.GetDataSourcePort()
}

func GetDataSourceDBName() string {
	return configInstance.GetDataSourceDBName()
}

func GetHttpOAuthAccessTokenSalt() string {
	return configInstance.GetHttpOAuthAccessTokenSalt()
}

func GetHttpOAuthAccessTokenValidTime() int64 {
	return configInstance.GetHttpOAuthAccessTokenValidTime()
}

func GetHttpOAuthRefreshTokenSalt() string {
	return configInstance.GetHttpOAuthRefreshTokenSalt()
}

func GetHttpOAuthRefreshTokenValidTime() int64 {
	return configInstance.GetHttpOAuthRefreshTokenValidTime()
}
