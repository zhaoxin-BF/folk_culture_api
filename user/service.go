/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/17 12:42 上午
 * @Description:
 */

package user

import (
	"fmt"
	"time"
)

//用户登陆model
type UserLogin struct {
	UserAccount    string       //账户
	UserPassword   string       //密码
}

//用户注册model
type UserRegister struct {
	UserName       string
	UserAccount    string
	UserPassword   string
	UserTel        string
}

//返回登陆数据model
type Resp struct {
	LoginStatus       int           //验证登陆是否成功，0成功  1失败
	UserInfo          interface{}
}

//用户登陆验证
func LoginLogic(form UserLogin) interface{}{
	var resp Resp
	//查找数据库，是否有该用户的信息
	userInfo, err := DBGetOneUser(form.UserAccount)

	if err != nil{
		resp.LoginStatus = 1
		resp.UserInfo = "非法用户，请输入正确的账户与密码！"
		return resp
	}
	//查看是否被封用户  3 为被封用户，不能登陆
	if userInfo.UserType == 3 {
		resp.LoginStatus = 1
		resp.UserInfo = "该用户因非法操作被封号，请联系超级管理员申请解封！"
		return resp
	}

	//比对数据，验证登陆
	if userInfo.UserAccount == form.UserAccount && (userInfo.UserPassword == form.UserPassword || userInfo.UserTel == form.UserPassword){
		resp.LoginStatus = 0
		resp.UserInfo = userInfo
		return resp
	}
	resp.LoginStatus = 1
	resp.UserInfo    = "用户登陆失败, 请检查账户与密码输入是否正确！"
	return resp
}

//用户注册
func RegisterUserLogic(form UsersTable) interface{}{
	if form.UserAccount == "" || form.UserPassword == "" || form.UserName == "" || form.UserTel == ""{
		return "必须填写好用户名、账户、密码、电话！"
	}
	form.UserType     = 0
	form.UserContext  = "普通用户"
	form.CreateTime   = time.Now().Unix()                                 //添加时间戳
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
//更新用户类型 0/普通用户 1/管理员 2/超级管理员/ 3封号中 /4删除用户
func UpdateUserTypeLogic(userId, userType int) interface{}{

	if userType == 4 {     //删除资源
		err := DBDeleteUser(userId)
		if nil != err {
			return "删除用户出错！"
		}
	}else {              //修改资源status
		err := DBUpdateUserType(userId,userType)
		if err != nil {
			return "修改用户类型出错！"
		}
	}
	return "修改用户类型成功！"
}
//删除一个用户信息
func DeleteUserLogic(userId int) interface{}{
	err := DBDeleteUser(userId)
	if err != nil{
		return "删除用户信息失败！"
	}
	return "删除用户信息成功！"
}