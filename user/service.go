/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/17 12:42 上午
 * @Description:
 */

package user

import "fmt"

//用户登陆model
type UserLogin struct {
	UserAccount    string       //账户
	UserPassword   string       //密码
}

//用户登陆验证
func LoginLogic(form UserLogin) interface{}{
	//查找数据库，是否有该用户的信息
	userInfo, err := DBGetOneUser(form.UserAccount)

	if err != nil{
		return "非法用户，请输入正确的账户与密码！"
	}
	if userInfo.UserAccount == form.UserAccount && (userInfo.UserPassword == form.UserPassword || userInfo.UserTel == form.UserPassword){
		return userInfo
	}
	return "用户登陆失败, 请检查账户与密码输入是否正确！"
}

//用户注册
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


//获取一个用户资料
func GetOneUserLogic(account string) interface{}{
	userInfo, err := DBGetOneUser(account)
	if err != nil {
		fmt.Println(err)
		return "获取用户信息失败！"
	}
	return userInfo
}
//获取所有用户资料
func GetAllUserLogic() interface{}{
	usersInfo, err := DBGetAllUser();
	if err != nil {
		return "获取所有用户信息失败！"
	}
	return usersInfo
}
//更新一个用户资料
func UpdateUserLogic(form UsersTable) interface{}{
	err := DBUpdateUser(&form)
	if err != nil {
		return "用户信息修改失败！"
	}
	return "用户信息修改成功！"
}
//删除一个用户信息
func DeleteUserLogic(account string) interface{}{
	err := DBDeleteUser(account)
	if err != nil{
		return "删除用户信息失败！"
	}
	return "删除用户信息成功！"
}