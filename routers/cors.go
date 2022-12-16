/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/9 2:23 下午
 * @Description:设置前端跨域处理，允许前端请求的访问
 */

package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//设置跨域处理函数
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		//可做请求次数入库记录

		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			//log.Info("---------------------------------------------------请求前端域为: ", origin)
			c.Header("Access-Control-Allow-Origin", origin) //设置全部跨域

			//服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//header的类型
			c.Header("Access-Control-Allow-Headers",
				"Authorization, Content-Length, X-CSRF-Token, Token, user_name, remote_user"+
					"session,X_Requested_With,Accept, Origin, Host, Connection,"+
					" Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive,"+
					" User-Agent, X-Requested-With, If-Modified-Since,"+
					" Cache-Control, Content-Type, Pragma")
			//允许跨域设置，可以返回其他子段,跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Expose-Headers",
				"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, remote_user,"+
					"Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")

			c.Header("Access-Control-Max-Age", "172800") // 缓存请求信息 单位为秒
			//必须设置为true
			c.Header("Access-Control-Allow-Credentials", "true") // 跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")            // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求,继续处理后续的请求
		c.Next()
	}
}

// cors middleware
//func CORS() gin.HandlerFunc {
//	config := cors.Config{
//		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
//		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "remote_user", "Token"},
//		AllowCredentials: true,
//		AllowOrigins:     []string{"http://localhost:3000"},
//		MaxAge:           12 * time.Hour,
//	}
//	return cors.New(config)
//}
