package oauth

import "frame/entity"

type User struct {
	entity.BaseEntity
	Id       int
	UserName string
	Password string
}
