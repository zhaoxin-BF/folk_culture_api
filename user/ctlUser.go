/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/17 12:42 上午
 * @Description:
 */

package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//注册
func RegisterUser(c *gin.Context){
	var form UsersTable

	//post 参数绑定
	if c.ShouldBind(&form) != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功 success!",
			"DataSet": "form表单提交参数错误！,请输入正确查询参数！",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet": RegisterUserLogic(form),
	})
}

//验证登陆一个用户
func Login(c *gin.Context){
	var form UserLogin

	//post 参数绑定
	if c.ShouldBind(&form) != nil {
		c.JSON(1, gin.H{
			"message": "请求失败 faild!",
			"DataSet": "form表单提交参数错误！,请输入正确的用户登陆信息！",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet":LoginLogic(form),
	})
}

//获取一个用户
func GetOneUser(c *gin.Context){
	userAccount := c.Query("account")
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet":GetOneUserLogic(userAccount),
	})
}

//获取全部用户
func GetAllUser(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet": GetAllUserLogic(),
	})
}

func UpdateUser(c *gin.Context){
	var form UsersTable

	//post 参数绑定
	if c.ShouldBind(&form) != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功 success!",
			"DataSet": "form表单提交参数错误！,请输入正确查询参数！",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet": UpdateUserLogic(form),
	})
}
//更新用户类型 0/普通用户 1/管理员 2/超级管理员/ 3封号中
func UpdateUserType(c *gin.Context){
	suser_id := c.Query("user_id")
	stype  := c.Query("type")

	//post 参数绑定
	userId, _ := strconv.Atoi(suser_id)
	userType , _ := strconv.Atoi(stype)

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet": UpdateUserTypeLogic(userId,userType),
	})
}

func DeleteUser(c *gin.Context){
	suser_id := c.Query("user_id")
	userId, _ := strconv.Atoi(suser_id)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet": DeleteUserLogic(userId),
	})
}