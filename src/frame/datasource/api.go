package datasource

import (
	"database/sql"
	"fmt"
	"frame/config"
	"frame/log"
)

var dbInstance *sql.DB

func Init() {
	errMsg := ""
	if !config.Exist(ConfigKeyDataBaseUserName) {
		errMsg = fmt.Sprint("frame.datasource.Init 缺少配置:[", ConfigKeyDataBaseUserName, "]")
	}
	if !config.Exist(ConfigKeyDataBasePassword) {
		errMsg = fmt.Sprint("frame.datasource.Init 缺少配置:[", ConfigKeyDataBaseUserName, "]")
	}
	if !config.Exist(ConfigKeyDataBaseHost) {
		errMsg = fmt.Sprint("frame.datasource.Init 缺少配置:[", ConfigKeyDataBaseUserName, "]")
	}
	if !config.Exist(ConfigKeyDataBasePort) {
		errMsg = fmt.Sprint("frame.datasource.Init 缺少配置:[", ConfigKeyDataBaseUserName, "]")
	}
	if !config.Exist(ConfigKeyDataBaseDbName) {
		errMsg = fmt.Sprint("frame.datasource.Init 缺少配置:[", ConfigKeyDataBaseUserName, "]")
	}
	if !config.Exist(ConfigKeyDataBaseDriver) {
		errMsg = fmt.Sprint("frame.datasource.Init 缺少配置:[", ConfigKeyDataBaseUserName, "]")
	}
	if errMsg != "" {
		log.Error(errMsg)
		panic(errMsg)
	}

	username := config.GetString(ConfigKeyDataBaseUserName)
	password := config.GetString(ConfigKeyDataBasePassword)
	host := config.GetString(ConfigKeyDataBaseHost)
	port := config.GetInt(ConfigKeyDataBasePort)
	dbName := config.Get(ConfigKeyDataBaseDbName)
	driver := config.GetString(ConfigKeyDataBaseDriver)

	url := fmt.Sprint(username, ":", password, "@tcp(", host, ":", port, ")/", dbName)
	db, err := sql.Open(driver, url)
	if err != nil {
		log.Error("frame.datasource.Init 数据库连接失败,错误信息:[", err, "]")
		return
	}
	dbInstance = db
	log.Info("数据库连接成功")

}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return dbInstance.Exec(query, args...)
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return dbInstance.Query(query, args...)
}
