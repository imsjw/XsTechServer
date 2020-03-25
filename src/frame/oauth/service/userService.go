package service

import (
	"oauth/dao"
	"oauth/entity"
)

func FindUserByUserNameAndPassword(username string, password string) *entity.User {
	return dao.SelectUserByUserNameAndPassword(username, password)
}
