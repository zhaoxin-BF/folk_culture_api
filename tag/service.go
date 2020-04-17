/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/17 12:42 上午
 * @Description:
 */

package tag

import "fmt"

func AddTagLogic(form TagsTable) interface{}{
	if form.TagName == ""{
		return "必须填写好分类名称！"
	}

	err := DBCreateTag(&form)
	if err != nil {
		return "用户注册失败，已存在该用户，请重新注册！"
	}
	return "用户注册成功！"
}

func GetOneTagLogic(tagId int) interface{}{
	userInfo, err := DBGetOneTag(tagId)
	if err != nil {
		fmt.Println(err)
		return "获取标签信息失败！"
	}
	return userInfo
}

func GetAllTagLogic() interface{}{
	usersInfo, err := DBGetAllTag();
	if err != nil {
		return "获取所有标签信息失败！"
	}
	return usersInfo
}

func UpdateTagLogic(form TagsTable) interface{}{
	err := DBUpdateTag(&form)
	if err != nil {
		return "标签信息修改失败！"
	}
	return "标签信息修改成功！"
}

func DeleteTagLogic(tagId int) interface{}{
	err := DBDeleteTag(tagId)
	if err != nil{
		return "删除标签信息失败！"
	}
	return "删除标签信息成功！"
}