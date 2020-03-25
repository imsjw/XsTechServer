package service

import (
	"frame/oauth/dao"
	"frame/oauth/entity"
)

func FindUserByUserNameAndPassword(username string, password string) *entity.User {
	return dao.SelectUserByUserNameAndPassword(username, password)
}
