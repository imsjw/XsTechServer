package oauth

import (
	"database/sql"
	"fmt"
	"frame/datasource"
)

var authDao = new(struct {
	datasource.BaseDao
})

func init() {
	authDao.TableName = "oauth_auth"
	authDao.Columns = "id,user_id,client,access_token,access_token_expires_time,refresh_token,refresh_token_expires_time,create_time,create_user,update_time,update_user"
}

func DaoInsertAuth(oauth *Auth) {
	kv := make(map[string]interface{})

	if oauth.Id != 0 {
		kv["id"] = oauth.Id
	}
	if oauth.UserId != 0 {
		kv["user_id"] = oauth.UserId
	}
	if oauth.Client != "" {
		kv["client"] = oauth.Client
	}
	if oauth.AccessToken != "" {
		kv["access_token"] = oauth.AccessToken
	}
	if oauth.AccessTokenExpiresTime != 0 {
		kv["access_token_expires_time"] = oauth.AccessTokenExpiresTime
	}
	if oauth.RefreshToken != "" {
		kv["refresh_token"] = oauth.RefreshToken
	}
	if oauth.RefreshTokenExpiresTime != 0 {
		kv["refresh_token_expires_time"] = oauth.RefreshTokenExpiresTime
	}
	if oauth.CreateTime != 0 {
		kv["create_time"] = oauth.CreateTime
	}
	if oauth.CreateUser != 0 {
		kv["create_user"] = oauth.CreateUser
	}
	if oauth.UpdateTime != 0 {
		kv["update_time"] = oauth.UpdateTime
	}
	if oauth.UpdateUser != 0 {
		kv["update_user"] = oauth.UpdateUser
	}

	sqlColumNames := ""
	rpStr := ""
	sqlColumValues := make([]interface{}, 0)

	for k, v := range kv {
		sqlColumNames += k
		sqlColumNames += ","
		rpStr += "?,"
		sqlColumValues = append(sqlColumValues, v)
	}

	sqlColumNames = sqlColumNames[:len(sqlColumNames)-1]
	rpStr = rpStr[:len(rpStr)-1]

	sql := fmt.Sprint("insert into ", authDao.TableName, " (", sqlColumNames+") values (", rpStr, ")")

	_, err := datasource.Exec(sql, sqlColumValues...)
	if err != nil {
		panic(err)
	}
}

func DaoDeleteAuthByUserIdAndClient(userId int, client string) {
	sql := fmt.Sprint("delete from ", authDao.TableName, " where user_id = ? and client = ?")
	_, err := datasource.Exec(sql, userId, client)
	if err != nil {
		panic(err)
	}
}

func DaoSelectAuthByUserIdAndClient(userId int, client string) *Auth {
	sql := fmt.Sprint("select ", authDao.Columns, " from ", authDao.TableName, " where user_id = ? and client = ?")
	rows, err := datasource.Query(sql, userId, client)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		return daoRowsToAuth(rows)
	}

	return nil
}

func DaoSelectOauthByAccessToken(accessToken string) *Auth {
	sql := fmt.Sprint("select ", authDao.Columns, " from ", authDao.TableName, " where access_token = ?")
	rows, err := datasource.Query(sql, accessToken)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		return daoRowsToAuth(rows)
	}
	return nil
}

func daoRowsToAuth(rows *sql.Rows) *Auth {
	var auth Auth
	err := rows.Scan(&auth.Id, &auth.UserId, &auth.Client, &auth.AccessToken, &auth.AccessTokenExpiresTime,
		&auth.RefreshToken, &auth.RefreshTokenExpiresTime, &auth.CreateTime, &auth.CreateUser, &auth.UpdateTime, &auth.UpdateUser)
	if err != nil {
		panic(err)
	}
	return &auth
}
