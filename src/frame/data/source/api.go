package source

import (
	"common/src/config"
	"database/sql"
	"fmt"
	"log"
)

var dbInstance *sql.DB

func Init() {
	if config.GetDataSourceUserName() == "" || config.GetDataSourcePassword() == "" || config.GetDataSourceHost() == "" || config.GetDataSourcePort() == 0 || config.GetDataSourceDBName() == "" {
		return
	}

	url := fmt.Sprint(config.GetDataSourceUserName(), ":", config.GetDataSourcePassword(), "@tcp(", config.GetDataSourceHost(), ":", config.GetDataSourcePort(), ")/", config.GetDataSourceDBName())
	db, err := sql.Open(config.GetDataSourceDriverName(), url)
	if err != nil {
		log.Println("data_source.Init error [", err, "]")
		return
	}
	dbInstance = db
	log.Println("data_source.Init success [database connect success]")

}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return dbInstance.Exec(query, args...)
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return dbInstance.Query(query, args...)
}
