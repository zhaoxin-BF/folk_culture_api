/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/7 2:23 下午
 * @Description:路由模块，设置跨域，设置路由组，由main.go 调用
 */

package routers

import (
	"folk_culture_api/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()

	//设置跨域中间件
	r.Use(Cors())

	//mikasa_ufs_api路由设置 第一版本v1
	v1Group := r.Group("v1")
	{
		//1、资源查询路由组—search
		schGroup := v1Group.Group("/search")
		{
			schGroup.GET("/detail",IndexData)              //一：资源详情查询接口
			schGroup.GET("/capacity",IndexData)              //二：资源容量查询接口
		}
	}

	//打印日志
	log.Info("routers.go","路由设置完毕!!!>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>项目启动成功!!!>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.")
	return r
}

func IndexData(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功 success!",
		"DataSet": nil,
	})
}