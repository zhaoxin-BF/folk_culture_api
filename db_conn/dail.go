/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/9 11:42 上午
 * @Description:初始化数据库连接，建立数据库连接池
 */

package db_conn

import (
	"folk_culture_api/log"
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

var (
	Database    *gorm.DB                                     //map[type+regionId].*Database
	once        sync.Once
)


//返回数据库连接
func GetUfsDatabase() *gorm.DB{
	once.Do(func() {
		dial()
	})
	return Database
}


//建立数据库初始化连接
func dial(){
	var dbUrl = "root:Zhaoxin..521@tcp(39.96.179.159:3306)/folk_culture_system?charset=utf8mb4"
	log.Info("---------------------初始化数据库连接中............")
	conn, err := gorm.Open("mysql", dbUrl)
	if err != nil{                         //检查连接池创建是否出现问题
		log.Error(err.Error())
	}
	err = conn.DB().Ping()
	if  err != nil {
		log.Error(err.Error())             //检查连接池内连接是否可用
	}
	conn.DB().SetMaxIdleConns(10)                   //设置最大空闲连接数
	conn.DB().SetMaxOpenConns(500)                  //设置最大打开连接数
	conn.DB().SetConnMaxLifetime(3 * time.Minute)      //设置连接超时，时间为3分钟

	Database = conn
	log.Info("---------------------数据库连接成功！---------------")
}