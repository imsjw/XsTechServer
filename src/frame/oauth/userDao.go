package oauth

import (
	"frame/datasource"
)

func DaoSelectUserByUserNameAndPassword(username string, password string) *User {
	sql := "select id,username from user where username = ? and password = ?"
	rows, err := datasource.Query(sql, username, password)
	if err != nil {
		panic(err)
	}

	var user User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.UserName)
		if err != nil {
			panic(err)
		}
		return &user
	}

	return nil
}
