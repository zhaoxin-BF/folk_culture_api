/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/17 12:41 上午
 * @Description:
 */

package res

import "fmt"

func AddResLogic(form ResourceTable) interface{}{
	//if form.UserAccount == "" || form.UserPassword == "" || form.UserName == "" || form.UserTel == ""{
	//	return "必须填写好用户名、账户、密码、电话！"
	//}

	err := DBCreateRes(&form)
	if err != nil {
		return "资源添加失败！"
	}
	return "资源添加成功！"
}

func GetOneResLogic(resId int) interface{}{
	resInfo, err := DBGetOneRes(resId)
	if err != nil {
		fmt.Println(err)
		return "获取资源信息失败！"
	}
	return resInfo
}

func GetAllResLogic() interface{}{
	resInfos, err := DBGetAllRes();
	if err != nil {
		return "获取所有资源信息失败！"
	}
	return resInfos
}

func UpdateResLogic(form ResourceTable) interface{}{
	err := DBUpdateRes(&form)
	if err != nil {
		return "资源信息修改失败！"
	}
	return "资源信息修改成功！"
}

func DeleteResLogic(resId int) interface{}{
	err := DBDeleteRes(resId)
	if err != nil{
		return "删除资源信息失败！"
	}
	return "删除资源信息成功！"
}