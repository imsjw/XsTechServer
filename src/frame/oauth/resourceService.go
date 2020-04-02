package oauth

func serviceExistResourceByUserId(userId int, url string, method string) bool {
	roldIds := daoSelectRoleIdsByUserId(userId)
	if len(roldIds) <= 0 {
		return false
	}

	resource := daoSelectResourceByUrlAndMethod(url, method)
	if resource == nil {
		return false
	}

	resources := daoSelectResourceIdsInRoleIds(roldIds...)
	if len(resources) <= 0 {
		return false
	}

	for _, v := range resources {
		if v == resource.Id {
			return true
		}
	}

	return false
}
