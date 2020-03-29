package oauth

import "frame/entity"

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
