package config

type DefaultConfig struct {
	Http struct {
		Server struct {
			Port int
		}
		OAuth struct {
			Enable   bool
			Password struct {
				Salt string
			}
			AccessToken struct {
				Salt      string
				ValidTime int64
			}
			RefreshToken struct {
				Salt      string
				ValidTime int64
			}
		}
	}
	DataSource struct {
		DriverName string
		UserName   string
		Password   string
		Host       string
		Port       int
		DBName     string
	}
}

func (This *DefaultConfig) GetHttpServerPort() int {
	return This.Http.Server.Port
}

func (This *DefaultConfig) GetHttpOAuthEnable() bool {
	return This.Http.OAuth.Enable
}

func (This *DefaultConfig) GetHttpOAuthPasswordSalt() string {
	return This.Http.OAuth.Password.Salt
}

func (This *DefaultConfig) GetDataSourceDriverName() string {
	return This.DataSource.DriverName
}

func (This *DefaultConfig) GetDataSourceUserName() string {
	return This.DataSource.UserName
}
func (This *DefaultConfig) GetDataSourcePassword() string {
	return This.DataSource.Password
}
func (This *DefaultConfig) GetDataSourceHost() string {
	return This.DataSource.Host
}
func (This *DefaultConfig) GetDataSourceDBName() string {
	return This.DataSource.DBName
}

func (This *DefaultConfig) GetDataSourcePort() int {
	return This.DataSource.Port
}

func (This *DefaultConfig) GetHttpOAuthAccessTokenSalt() string {
	return This.Http.OAuth.AccessToken.Salt
}

func (This *DefaultConfig) GetHttpOAuthAccessTokenValidTime() int64 {
	return This.Http.OAuth.AccessToken.ValidTime
}

func (This *DefaultConfig) GetHttpOAuthRefreshTokenSalt() string {
	return This.Http.OAuth.RefreshToken.Salt
}

func (This *DefaultConfig) GetHttpOAuthRefreshTokenValidTime() int64 {
	return This.Http.OAuth.RefreshToken.ValidTime
}
