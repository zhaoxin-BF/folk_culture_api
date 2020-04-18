/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/7 2:23 下午
 * @Description:路由模块，设置跨域，设置路由组，由main.go 调用
 */

package routers

import (
	"folk_culture_api/log"
	res "folk_culture_api/resource"
	"folk_culture_api/tag"
	"folk_culture_api/user"
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
		//1、用户查询路由组—user
		userGroup := v1Group.Group("/user")
		{
			userGroup.POST("/login",user.Login)
			userGroup.GET("/getOne",user.GetOneUser)                //获得一个用户,id
			userGroup.GET("/getAll",user.GetAllUser)                //获得全部用户
			userGroup.POST("/register",user.RegisterUser)           //注册一个用户
			userGroup.GET("/delete",user.DeleteUser)                //删除一个用户,id
			userGroup.POST("/update",user.UpdateUser)               //更新一个用户,id
		}

		//2、标签查询路由组—tag
		tagGroup := v1Group.Group("/tag")
		{
			tagGroup.GET("/getOne",tag.GetOneTag)                    //获得标签信息,id
			tagGroup.GET("/getAll",tag.GetAllTag)                    //获得全部标签
			tagGroup.POST("/add",tag.AddTag)                         //添加一个标签
			tagGroup.GET("/delete",tag.DeleteTag)                    //删除一个标签,id
			tagGroup.POST("/update",tag.UpdateTag)                   //更新一个标签,id
		}

		//3、资源查询路由组—resource
		resGroup := v1Group.Group("/res")
		{
			resGroup.GET("/getOne",res.GetOneRes)                    //获得一条资源,id
			resGroup.GET("/searchRes",res.GetResByResName)           //获得资源模糊查询，resName
			resGroup.GET("/getTagRes",res.GetResByTagId)             //获取同一类型的资源，tagTd
			resGroup.GET("/getAll",res.GetAllRes)                    //获得全部资源
			resGroup.POST("/add",res.AddRes)                         //添加一个资源
			resGroup.GET("/delete",res.DeleteRes)                    //删除一个资源,id
			resGroup.POST("/update",res.UpdateRes)                   //更新一个资源,id
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