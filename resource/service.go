/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/17 12:41 上午
 * @Description:
 */

package res

import (
	"fmt"
	"time"
)
//前端添加资源，绑定数据model
type FormResTable struct {
	ResourceId           int
	ResourceName         string
	TagName              string
	Description          string
	Author               string
	Time                 string      //年代
	Nation               string      //民族
	Region               string      //地域
	Url                  string
	ResourceContext      string      //备注信息

	UploadId             int         //user_id
	UploadUser           string      //user_name
}


func AddResLogic(form FormResTable) interface{}{
	//创建新的插入资源信息变量
	var resInfo ResourceTable

	//获取到前端传入的数据
	resInfo.ResourceName    = form.ResourceName
	resInfo.TagName         = form.TagName
	resInfo.Description     = form.Description
	resInfo.Author          = form.Author
	resInfo.Time            = form.Time
	resInfo.Nation          = form.Nation
	resInfo.Region          = form.Region
	resInfo.Url             = form.Url
	resInfo.ResourceContext = form.ResourceContext           //具体内容
	resInfo.UploadId        = form.UploadId
	resInfo.UploadUser      = form.UploadUser

	//获得时间信息


	//后端添加时间信息，状态信息
	resInfo.Status          = 1
	resInfo.CreateTime      = time.Now().Unix()                                 //添加时间戳
	resInfo.ScreateTime     = time.Now().Format("2006-01-02 15:04:05")   //添加创建时间 string


	err := DBCreateRes(&resInfo)
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

//获得所有通过审核的资源的信息， user
func GetAllResLogic() interface{}{
	status := 0
	resInfos, err := DBGetAllRes(status);
	if err != nil {
		return nil
	}
	return resInfos
}

func MGetAllResLogic() interface{} {
	allInfos, err := DBMGetAllRes();
	resInfos := make(map[int][]ResourceTable)

	//分类，通过0，审核中1，没通过2
	for _, val := range allInfos {
		resInfos[val.Status] = append(resInfos[val.Status], val)
	}
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

//修改资源状态 0/1/2          3
func UpdateResStatusLogic(check_name string,resId ,status int) interface{}{

	if status == 3 {     //删除资源
		err := DBDeleteRes(resId)
		if nil != err {
			return "删除资源信息出错！"
		}
	}else {              //修改资源status
		err := DBUpdateResStatus(check_name,resId, status)
		if err != nil {
			return "修改资源状态出错！"
		}
	}
	return "修改资源状态成功！"
}