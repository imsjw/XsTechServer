package oauth

import (
	"database/sql"
	"fmt"
	"frame/datasource"
)

var userDao = new(struct {
	datasource.BaseDao
})

func init() {
	userDao.TableName = "user"
	userDao.Columns = "id,username,create_time,create_user,update_time,update_user"
}

func daoSelectUserByUserNameAndPassword(username string, password string) *User {
	sql := fmt.Sprint("select ", userDao.Columns, " from ", userDao.TableName, " where username = ? and password = ?")
	rows, err := datasource.Query(sql, username, password)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		return daoRowsToUser(rows)
	}

	return nil
}

func daoRowsToUser(rows *sql.Rows) *User {
	var user User
	err := rows.Scan(&user.Id, &user.UserName, &user.CreateTime, &user.CreateUser, &user.UpdateTime, &user.UpdateUser)
	if err != nil {
		panic(err)
	}
	return &user
}
