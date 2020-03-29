package oauth

import "frame/datasource"

var userRoleDao = new(struct {
	datasource.BaseDao
})

func init() {
	userRoleDao.TableName = "oauth_user_role"
	userRoleDao.Columns = "id,create_time,create_user,update_time,update_user"
}
