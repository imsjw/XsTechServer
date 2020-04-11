package oauth

func ServiceFindUserByUserNameAndPassword(username string, password string) *User {
	return daoSelectUserByUserNameAndPassword(username, password)
}
