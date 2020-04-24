/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/23 3:13 下午
 * @Description:
 */

package auth_demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//model
type Role struct {
	Menus []string
}

//添加一个资源
func AddRole(c *gin.Context){
	cookie,_ := c.Cookie("user");
	fmt.Println(cookie)
	var form Role

	//post 参数绑定
	if c.ShouldBind(&form) != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "form表单提交参数错误！,请输入正确查询参数！",
			"DataSet": nil,
		})
	}
	fmt.Println(form)

	for _, val := range form.Menus{
		fmt.Println(val)
	}
}