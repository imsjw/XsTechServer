package oauth

import (
	"crypto/md5"
	"encoding/hex"
	"frame/entity"
	"time"
)

func ServicePassworMethodAuthorize(username string, rawPassword string, client string) interface{} {
	encodePassword := ServicePasswordEncryption(rawPassword)
	user := ServiceFindUserByUserNameAndPassword(username, encodePassword)
	if user == nil {
		return ResultUserOrPasswordError
	}
	auth := daoSelectAuthByUserIdAndClient(user.Id, client)
	res := new(struct {
		UserId                 int
		Client                 string
		AccessToken            string
		AccessTokenExpiresTime int64
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

func ServicePasswordEncryption(rawPassword string) string {
	salt := configPasswordSalt
	h := md5.New()
	h.Write([]byte(rawPassword))
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum(nil))
}

func ServiceGetAuthByAccessToken(accessToken string) *Auth {
	auth := daoSelectOauthByAccessToken(accessToken)
	return auth
}
