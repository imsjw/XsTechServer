package oauth

import "frame/entity"

type User struct {
	entity.BaseEntity
	Id       int
	UserName string
	Password string
}

type Auth struct {
	entity.BaseEntity
	Id                      int
	UserId                  int
	Client                  string
	AccessToken             string
	AccessTokenExpiresTime  int64
	RefreshToken            string
	RefreshTokenExpiresTime int64
}

type Resource struct {
	entity.BaseEntity
	Id     int
	Method string
	Url    string
}

type UserRole struct {
	entity.BaseEntity
	Id     int
	UserId int
	RoleId int
}
