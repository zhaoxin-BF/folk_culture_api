/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2022/12/1 11:03
 * @Description:
 */

package user

import (
	"fmt"
	"folk_culture_api/db_conn"
	"testing"
)

func TestDBGetOneUser(t *testing.T) {
	db_conn.InitMySQL()
	userInfo, err := DBGetOneUser("boreas")
	if err != nil {
		fmt.Println("非法用户，请输入正确的账户与密码！")
	}
	fmt.Println(userInfo)
}

func TestDBGetAllUser(t *testing.T) {
	db_conn.InitMySQL()
	userInfos, err := DBGetAllUser()
	if err != nil {
		fmt.Println("非法用户，请输入正确的账户与密码！")
	}
	fmt.Println(userInfos)
}
