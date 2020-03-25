package config

type Config interface {
	GetHttpServerPort() int

	GetHttpOAuthEnable() bool
	GetHttpOAuthPasswordSalt() string
	GetHttpOAuthAccessTokenSalt() string
	GetHttpOAuthAccessTokenValidTime() int64
	GetHttpOAuthRefreshTokenSalt() string
	GetHttpOAuthRefreshTokenValidTime() int64

	GetDataSourceDriverName() string
	GetDataSourceUserName() string
	GetDataSourcePassword() string
	GetDataSourceHost() string
	GetDataSourcePort() int
	GetDataSourceDBName() string
}
