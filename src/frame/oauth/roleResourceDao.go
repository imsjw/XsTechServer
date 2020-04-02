package oauth

import (
	"fmt"
	"frame/datasource"
)

var roleResourceDao = new(struct {
	datasource.BaseDao
})

func init() {
	userRoleDao.TableName = "oauth_role_resource"
	userRoleDao.Columns = "id,create_time,create_user,update_time,update_user"
}

func daoSelectResourceIdsInRoleIds(roldIds ...int) []int {
	sql := fmt.Sprint("select * from ", roleResourceDao.TableName, " where role_id in (%s)")
	inStatus := ""
	pSize := len(roldIds)
	for i := 0; i < pSize; i++ {
		if i == 0 {
			inStatus += "?"
		} else {
			inStatus += ",?"
		}
	}
	sql = fmt.Sprintf(sql, inStatus)
	rows, err := datasource.Query(sql, roldIds)
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
