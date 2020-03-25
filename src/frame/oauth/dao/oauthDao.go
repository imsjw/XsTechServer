package dao

import (
	"fmt"
	"frame/data/source"
	"frame/oauth/entity"
)

func SelectOauthByUserIdAndClient(userId int, client string) *entity.Oauth {
	sql := "select " +
		"id,user_id,client,access_token,access_token_expires_time,refresh_token,refresh_token_expires_time," +
		"create_time,create_user,update_time,update_user " +
		"from oauth where user_id = ? and client = ?"
	rows, err := source.Query(sql, userId, client)
	if err != nil {
		panic(err)
	}

	var oauth entity.Oauth
	for rows.Next() {
		err := rows.Scan(&oauth.Id, &oauth.UserId, &oauth.Client, &oauth.AccessToken, &oauth.AccessTokenExpiresTime,
			&oauth.RefreshToken, &oauth.RefreshTokenExpiresTime, &oauth.CreateTime, &oauth.CreateUser, &oauth.UpdateTime, &oauth.UpdateUser)
		if err != nil {
			panic(err)
		}
		return &oauth
	}

	return nil
}

func DeleteOauthByUserIdAndClient(userId int, client string) {
	sql := "delete from oauth where user_id = ? and client = ?"
	_, err := source.Exec(sql, userId, client)
	if err != nil {
		panic(err)
	}
}

func InsertOauth(oauth *entity.Oauth) {
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

	sql := fmt.Sprint("insert into oauth (", sqlColumNames+") values (", rpStr, ")")

	_, err := data_source.Exec(sql, sqlColumValues...)
	if err != nil {
		panic(err)
	}
}
