package oauth

import (
	"crypto/md5"
	"encoding/hex"
	"frame/entity"
	"time"
)

func servicePassworMethodAuthorize(username string, rawPassword string, client string) interface{} {
	encodePassword := servicePasswordEncryption(rawPassword)
	user := ServiceFindUserByUserNameAndPassword(username, encodePassword)
	if user == nil {
		return ResultUserOrPasswordError
	}
	auth := daoSelectAuthByUserIdAndClient(user.Id, client)
	res := new(struct {
		UserId                 int    `json:"userId"`
		Client                 string `json:"client"`
		AccessToken            string `json:"accessToken"`
		AccessTokenExpiresTime int64  `json:"accessTokenExpiresTime"`
	})

	//删除旧token
	daoDeleteAuthByUserIdAndClient(user.Id, client)
	//创建新的token
	auth = new(Auth)
	auth.UserId = user.Id
	auth.Client = client
	currTime := time.Now().Unix()
	auth.AccessTokenExpiresTime = currTime + configAccessTokenValidTime
	auth.RefreshTokenExpiresTime = currTime + configRefreshTokenValidTime
	auth.CreateTime = time.Now().Unix()
	auth.UpdateTime = time.Now().Unix()
	auth.CreateUser = user.Id
	auth.UpdateUser = user.Id

	accessToken := jwtHS256(auth, configAccessTokenSalt)
	refreshToken := jwtHS256(auth, configRefreshTokenSalt)
	auth.AccessToken = accessToken
	auth.RefreshToken = refreshToken

	daoInsertAuth(auth)

	res.UserId = user.Id
	res.Client = client
	res.AccessToken = accessToken
	res.AccessTokenExpiresTime = auth.AccessTokenExpiresTime
	return entity.NewSuccessResult(res)
}

func servicePasswordEncryption(rawPassword string) string {
	salt := configPasswordSalt
	h := md5.New()
	h.Write([]byte(rawPassword))
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum(nil))
}

func serviceGetAuthByAccessToken(accessToken string) *Auth {
	auth := daoSelectOauthByAccessToken(accessToken)
	return auth
}

/**
根据authId刷新token
*/
func serviceRefreshTokenById(auth *Auth) *Auth {
	currTime := time.Now().Unix()

	newAuth := new(Auth)
	newAuth.UserId = auth.UserId
	newAuth.Client = auth.Client
	newAuth.AccessTokenExpiresTime = currTime + configAccessTokenValidTime
	newAuth.RefreshTokenExpiresTime = currTime + configRefreshTokenValidTime
	newAuth.CreateTime = auth.CreateTime
	newAuth.UpdateTime = time.Now().Unix()
	newAuth.CreateUser = auth.CreateUser
	newAuth.UpdateUser = auth.UserId

	accessToken := jwtHS256(newAuth, configAccessTokenSalt)
	refreshToken := jwtHS256(newAuth, configRefreshTokenSalt)
	newAuth.AccessToken = accessToken
	newAuth.RefreshToken = refreshToken

	newAuth.Id = auth.Id

	daoRefreshTokenById(newAuth)
	return daoSelectOauthById(auth.Id)
}
