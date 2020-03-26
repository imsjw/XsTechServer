package dao

import (
	"frame/db"
	"frame/oauth/entity"
)

var userDao = new(struct {
	db.BaseDao
})

func init() {
	userDao.TableName = "user"
}

func SelectUserByUserNameAndPassword(username string, password string) *entity.User {
	sql := "select id,username from user where username = ? and password = ?"
	rows, err := db.Query(sql, username, password)
	if err != nil {
		panic(err)
	}

	var user entity.User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.UserName)
		if err != nil {
			panic(err)
		}
		return &user
	}

	return nil
}
