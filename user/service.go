/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/17 12:42 上午
 * @Description:
 */

package user

import "fmt"

func RegisterUserLogic(form UsersTable) interface{}{
	if form.UserAccount == "" || form.UserPassword == "" || form.UserName == "" || form.UserTel == ""{
		return "必须填写好用户名、账户、密码、电话！"
	}

	err := DBCreateUser(&form)
	if err != nil {
		return "用户注册失败，已存在该用户，请重新注册！"
	}
	return "用户注册成功！"
}

func GetOneUserLogic(account string) interface{}{
	userInfo, err := DBGetOneUser(account)
	if err != nil {
		fmt.Println(err)
		return "获取用户信息失败！"
	}
	return userInfo
}

func GetAllUserLogic() interface{}{
	usersInfo, err := DBGetAllUser();
	if err != nil {
		return "获取所有用户信息失败！"
	}
	return usersInfo
}

func UpdateUserLogic(form UsersTable) interface{}{
	err := DBUpdateUser(&form)
	if err != nil {
		return "用户信息修改失败！"
	}
	return "用户信息修改成功！"
}

func DeleteUserLogic(account string) interface{}{
	err := DBDeleteUser(account)
	if err != nil{
		return "删除用户信息失败！"
	}
	return "删除用户信息成功！"
}