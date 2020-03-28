package oauth

func FindUserByUserNameAndPassword(username string, password string) *User {
	return SelectUserByUserNameAndPassword(username, password)
}
