package oauth

func ServiceFindUserByUserNameAndPassword(username string, password string) *User {
	return DaoSelectUserByUserNameAndPassword(username, password)
}
