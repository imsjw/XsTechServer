package oauth

import (
	"database/sql"
	"fmt"
	"frame/datasource"
)

var resourceDao = new(struct {
	datasource.BaseDao
})

func init() {
	resourceDao.TableName = "oauth_resource"
	resourceDao.Columns = "id,client,method,url,create_time,create_user,update_time,update_user"
}

func SelectCount(p *Resource) int {
	sql := fmt.Sprint("select count(1) from ", resourceDao.TableName)
	sqlP := []interface{}{}
	var existP = false
	whereStr := ""
	if p.Id != 0 {
		sqlP = append(sqlP, p.Id)
		whereStr = fmt.Sprint(whereStr, " and id = ?")
		existP = true
	}
	if p.Client != "" {
		sqlP = append(sqlP, p.Client)
		whereStr = fmt.Sprint(whereStr, " and client = ?")
		existP = true
	}
	if p.Method != "" {
		sqlP = append(sqlP, p.Method)
		whereStr = fmt.Sprint(whereStr, " and method = ?")
		existP = true
	}
	if p.Url != "" {
		sqlP = append(sqlP, p.Url)
		whereStr = fmt.Sprint(whereStr, " and url = ?")
		existP = true
	}

	if existP {
		sql = fmt.Sprint(sql, " where ", sqlP[5:])
	}

	rows, err := datasource.Query(sql, sqlP...)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var count int
		err := rows.Scan(&count)
		if err != nil {
			panic(err)
		}
		return count
	}
	return 0
}

func daoRowsToResource(rows *sql.Rows) *Resource {
	var res Resource
	err := rows.Scan(&res.Id, &res.Client, &res.Method, &res.Url, &res.CreateTime, &res.CreateUser, &res.UpdateTime, &res.UpdateUser)
	if err != nil {
		panic(err)
	}
	return &res
}
