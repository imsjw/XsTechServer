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
	oauth := DaoSelectOauthByUserIdAndClient(user.Id, client)
	res := new(struct {
		UserId                 int
		Client                 string
		AccessToken            string
		AccessTokenExpiresTime int64
	})

	//删除旧token
	DaoDeleteOauthByUserIdAndClient(user.Id, client)
	//创建新的token
	oauth = new(Auth)
	oauth.UserId = user.Id
	oauth.Client = client
	currTime := time.Now().Unix()
	oauth.AccessTokenExpiresTime = currTime + configAccessTokenValidTime
	oauth.RefreshTokenExpiresTime = currTime + configRefreshTokenValidTime
	oauth.CreateTime = time.Now().Unix()
	oauth.UpdateTime = time.Now().Unix()
	oauth.CreateUser = user.Id
	oauth.UpdateUser = user.Id

	accessToken := JwtHS256(oauth, configAccessTokenSalt)
	refreshToken := JwtHS256(oauth, configRefreshTokenSalt)
	oauth.AccessToken = accessToken
	oauth.RefreshToken = refreshToken

	InsertOauth(oauth)

	res.UserId = user.Id
	res.Client = client
	res.AccessToken = accessToken
	res.AccessTokenExpiresTime = oauth.AccessTokenExpiresTime
	return entity.NewSuccessResult(res)
}

func ServicePasswordEncryption(rawPassword string) string {
	salt := configPasswordSalt
	h := md5.New()
	h.Write([]byte(rawPassword))
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum(nil))
}

func ServiceGetOauthByAccessToken(accessToken string) *Auth {
	oauth := DaoSelectOauthByAccessToken(accessToken)
	return oauth
}
