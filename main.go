/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/16 11:31 下午
 * @Description:
 */

package main

import (
	_ "github.com/go-sql-driver/mysql"       //注意⚠️不含该包会报错 sql: unknown driver "mysql"
	"folk_culture_api/db_conn"
	"folk_culture_api/log"
	"folk_culture_api/routers"
)

func main(){
	//一、启动日志模块配置
	log.InitLog()


	//三、启动数据库连接模块
	db_conn.InitMysql()

	//end、启动设置路由
	r := routers.SetupRouter()

	//end、启动设置启动端口
	r.Run(":6666")
}