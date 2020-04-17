/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/17 12:42 上午
 * @Description:
 */

package user

import "folk_culture_api/db_conn"

type UsersTable struct {
	UserId         int
	UserName       string
	UserAccount    string
	UserPassword   string
	UserTel        string
	UserType       int
	UserContext    string
}

// 设置UsersTable 的表名为`users_table`                            //⚠️非常值得注意：数据库表名与model的映射关系
func (UsersTable) TableName() string {
	return "users_table"
}


func GetUserInfo()  {

	return
}

func GetAllUserInfo() {
	return
}

func UpdateUser()  {
	return
}

func InsertUser() {
	var user UsersTable
	user.UserAccount  = "1111111"
	user.UserPassword = "1111111"
	user.UserName     = "赵亚娟"
	user.UserTel      = "13888888888"
	user.UserType     = 1
	user.UserContext  = ""
	db_conn.DB.Create(&user)
	return
}

func DeleteUser() {
	return
}