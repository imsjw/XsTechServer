package dao

import (
	"frame/data/source"
	"frame/oauth/entity"
)

var userDao = new(struct {
	source.BaseDao
})

func init() {
	userDao.TableName = "user"
}

func SelectUserByUserNameAndPassword(username string, password string) *entity.User {
	sql := "select id,username from user where username = ? and password = ?"
	rows, err := source.Query(sql, username, password)
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
