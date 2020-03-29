package oauth

import (
	"crypto/md5"
	"encoding/hex"
	sysEntity "frame/entity"
	"time"
)

func PassworMethodAuthorize(username string, rawPassword string, client string) interface{} {
	encodePassword := PasswordEncryption(rawPassword)
	user := ServiceFindUserByUserNameAndPassword(username, encodePassword)
	if user == nil {
		return ResultUserOrPasswordError
	}
	oauth := DaoSelectOauthByUserIdAndClient(user.Id, client)
	res := new(struct {
		sysEntity.BaseResult
		UserId                 int
		Client                 string
		AccessToken            string
		AccessTokenExpiresTime int64
	})

	//删除旧token
	DaoDeleteOauthByUserIdAndClient(user.Id, client)
	//创建新的token
	oauth = new(Oauth)
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

	res.Code = sysEntity.ResultCodeOk
	res.Msg = sysEntity.ResultMsgOk
	res.UserId = user.Id
	res.Client = client
	res.AccessToken = accessToken
	res.AccessTokenExpiresTime = oauth.AccessTokenExpiresTime
	return res
}

func PasswordEncryption(rawPassword string) string {
	salt := configPasswordSalt
	h := md5.New()
	h.Write([]byte(rawPassword))
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum(nil))
}
