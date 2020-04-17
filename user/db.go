/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/17 12:42 上午
 * @Description:
 */

package user

import (
	"folk_culture_api/db_conn"
)

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


//根据账户获取一个用户的详细信息
func DBGetOneUser(account string)(userInfo UsersTable, err error)  {
	err = db_conn.DB.Where("user_account = ?",account).First(&userInfo).Error;
	return
}

//获取所有用户的信息
func DBGetAllUser()(usersInfo []UsersTable, err error) {
	err = db_conn.DB.First(&usersInfo).Error;
	if err != nil {
		return nil, err
	}
	return
}

//更新用户的信息
func DBUpdateUser(userInfo *UsersTable)(err error){
	var user UsersTable
	err = db_conn.DB.Model(&user).Where("user_account = ?", userInfo.UserAccount).Updates(&userInfo).Error
	return
}

func DBCreateUser(userInfo *UsersTable) (err error){
	err = db_conn.DB.Create(&userInfo).Error
	return
}

func DBDeleteUser(account string) (err error){
	err = db_conn.DB.Where("user_account = ?",account).Delete(&UsersTable{}).Error
	return
}