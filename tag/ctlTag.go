/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/17 12:42 上午
 * @Description:
 */

package tag

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//注册
func AddTag(c *gin.Context){
	var form TagsTable

	//post 参数绑定
	if c.ShouldBind(&form) != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功 success!",
			"DataSet": "form表单提交参数错误！,请输入正确查询参数！",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet": AddTagLogic(form),
	})
}

//获取一个用户
func GetOneTag(c *gin.Context){
	stagId := c.Query("tag_id")
	tagId,_ := strconv.Atoi(stagId)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet":GetOneTagLogic(tagId),
	})
}

//获取全部用户
func GetAllTag(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet": GetAllTagLogic(),
	})
}

func UpdateTag(c *gin.Context){
	var form TagsTable

	//post 参数绑定
	if c.ShouldBind(&form) != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功 success!",
			"DataSet": "form表单提交参数错误！,请输入正确查询参数！",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet": UpdateTagLogic(form),
	})
}

func DeleteTag(c *gin.Context){
	stagId := c.Query("tag_id")
	tagId, _ := strconv.Atoi(stagId)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet": DeleteTagLogic(tagId),
	})
}