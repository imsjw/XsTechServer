package oauth

import (
	"fmt"
	"frame/datasource"
)

var userRoleDao = new(struct {
	datasource.BaseDao
})

func init() {
	userRoleDao.TableName = "oauth_user_role"
	userRoleDao.Columns = "id,create_time,create_user,update_time,update_user"
}

func daoSelectRoleIdsByUserId(userId int) []int {
	sql := fmt.Sprint("select role_id from ", userRoleDao.TableName, " where user_id = ?")
	rows, err := datasource.Query(sql, userId)
	if err != nil {
		panic(err)
	}

	res := []int{}

	for rows.Next() {
		var roleId int
		err := rows.Scan(&roleId)
		if err != nil {
			panic(err)
		}
		res = append(res, roleId)
	}

	return res
}
