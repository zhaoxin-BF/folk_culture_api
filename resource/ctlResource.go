/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/17 12:41 上午
 * @Description:
 */

package res

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加一个资源
func AddRes(c *gin.Context){
	var form ResourceTable

	//post 参数绑定
	if c.ShouldBind(&form) != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功 success!",
			"DataSet": "form表单提交参数错误！,请输入正确查询参数！",
		})
	}
	fmt.Println(form)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet": AddResLogic(form),
	})
}

//获取一个资源，根据res_id
func GetOneRes(c *gin.Context){
	sres_id := c.Query("res_id")
	resId, _ := strconv.Atoi(sres_id)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet":GetOneResLogic(resId),
	})
}

//根据resName 模糊查询资源信息
func GetResByResName(c *gin.Context){
	resName := c.Query("res_name")

	c.JSON(http.StatusOK, gin.H{
		"message":"请求成功 success",
		"DataSet":GetResByNameLogic(resName),
	})
}

//根据tagId 获取同一类型的资源信息
func GetResByTagId(c *gin.Context){
	stagId := c.Query("tag_id")
	tagId, _ := strconv.Atoi(stagId)
	c.JSON(http.StatusOK, gin.H{
		"message":"请求成功 success",
		"DataSet":GetResByTagIdLogic(tagId),
	})
}

//获取全部资源
func GetAllRes(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet": GetAllResLogic(),
	})
}

//更新一个资源
func UpdateRes(c *gin.Context){
	var form ResourceTable

	//post 参数绑定
	if c.ShouldBind(&form) != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功 success!",
			"DataSet": "form表单提交参数错误！,请输入正确查询参数！",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet": UpdateResLogic(form),
	})
}

//删除一个资源
func DeleteRes(c *gin.Context){
	sres_id := c.Query("res_id")
	resId, _ := strconv.Atoi(sres_id)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet": DeleteResLogic(resId),
	})
}