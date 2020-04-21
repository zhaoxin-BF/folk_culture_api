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

//根据资源id获取到一条数据
func GetOneResLogic(resId int) interface{}{
	resInfo, err := DBGetOneRes(resId)
	if err != nil {
		fmt.Println(err)
		return "获取资源信息失败！"
	}
	return resInfo
}

//根据资源名字resName，获取资源信息，模糊查询资源
func GetResByNameLogic(resName string) interface{} {
	resInfos, err := DBGetResByName(resName)
	if err != nil {
		return nil
	}

	if len(resInfos) == 0{
		return nil
	}
	return resInfos
}

//根据tag_id，获取同一类型资源信息
func GetResByTagIdLogic(tagId int) interface{} {
	resInfos, err := DBGetAllResById(tagId)
	if err != nil {
		return nil
	}

	if len(resInfos) == 0{
		return nil
	}
	return resInfos
}

//获得所有资源的信息
func GetAllResLogic() interface{}{
	resInfos, err := DBGetAllRes();
	if err != nil {
		return nil
	}
	return resInfos
}

//获得所有资源的信息
//func GetAllResByTagIdLogic() interface{}{
//	resInfos, err := DBGetAllRes();
//	if err != nil {
//		return "获取所有资源信息失败！"
//	}
//	return resInfos
//}

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