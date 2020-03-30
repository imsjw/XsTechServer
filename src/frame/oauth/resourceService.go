package oauth

func ServiceExistResourceByUserId(userId int) bool {
	roldIds := DaoSelectRoleIdsByUserId(userId)
	if len(roldIds) <= 0 {
		return false
	}
	
	return false
}
