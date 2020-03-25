package service

import (
	"crypto/md5"
	"encoding/hex"
	"frame/config"
	sysConstant "frame/constant"
	sysEntity "frame/entity"
	oauthConstant "oauth/constant"
	"oauth/dao"
	"oauth/entity"
	"oauth/util"
	"time"
)

func PassworMethodAuthorize(username string, rawPassword string, client string) interface{} {
	encodePassword := PasswordEncryption(rawPassword)
	user := FindUserByUserNameAndPassword(username, encodePassword)
	if user == nil {
		return oauthConstant.ResultUserOrPasswordError
	}
	oauth := dao.SelectOauthByUserIdAndClient(user.Id, client)
	res := new(struct {
		sysEntity.BaseResult
		UserId                 int
		Client                 string
		AccessToken            string
		AccessTokenExpiresTime int64
	})

	//删除旧token
	dao.DeleteOauthByUserIdAndClient(user.Id, client)
	//创建新的token
	oauth = new(entity.Oauth)
	oauth.UserId = user.Id
	oauth.Client = client
	currTime := time.Now().Unix()
	oauth.AccessTokenExpiresTime = currTime + config.GetHttpOAuthAccessTokenValidTime()
	oauth.RefreshTokenExpiresTime = currTime + config.GetHttpOAuthRefreshTokenValidTime()
	oauth.CreateTime = time.Now().Unix()
	oauth.UpdateTime = time.Now().Unix()
	oauth.CreateUser = user.Id
	oauth.UpdateUser = user.Id

	accessToken := util.JwtHS256(oauth, config.GetHttpOAuthAccessTokenSalt())
	refreshToken := util.JwtHS256(oauth, config.GetHttpOAuthRefreshTokenSalt())
	oauth.AccessToken = accessToken
	oauth.RefreshToken = refreshToken

	dao.InsertOauth(oauth)

	res.Code = sysConstant.ResultCodeOk
	res.Msg = sysConstant.ResultMsgOk
	res.UserId = user.Id
	res.Client = client
	res.AccessToken = accessToken
	res.AccessTokenExpiresTime = oauth.AccessTokenExpiresTime
	return res
}

func PasswordEncryption(rawPassword string) string {
	salt := config.GetHttpOAuthPasswordSalt()
	h := md5.New()
	h.Write([]byte(rawPassword))
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum(nil))
}
