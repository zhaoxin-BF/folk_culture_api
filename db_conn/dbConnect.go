/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/17 12:55 上午
 * @Description:
 */

package db_conn

import (
	"folk_culture_api/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:Zhaoxin..521@tcp(39.96.179.159:3306)/folk_culture_system?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	log.Info("mysql.go", "----------------------------------数据库链接成功！--------------------------------------")
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
