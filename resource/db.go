/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/17 12:42 上午
 * @Description:
 */

package res

import (
	"folk_culture_api/db_conn"
)

type ResourceTable struct {
	ResourceId           int
	ResourceName         string
	TagId                int
	TagName              string
	Description          string
	Author               string
	Time                 string      //年代
	Nation               string      //民族
	Region               string      //地域
	Copyright            string
	Url                  string
	Status               int         //资源状态 0 表示审核通过， 1 表示审核中 2 表示未通过审核

	CreateTime           int         //时间戳
	ScreateTime          string      //string时间
	UpdateTime           int         //时间戳
	SupdateTime          string      //string时间

	ResourceContext      string      //备注信息

	UploadId             int         //user_id
	UploadUser           string      //user_name

	CheckId              int         //user_id
	CheckName            string      //user_name
}

// 设置ResourceTable 的表名为`resources_table`                            //⚠️非常值得注意：数据库表名与model的映射关系
func (ResourceTable) TableName() string {
	return "resources_table"
}


//根据resourceId获取一个资源的详细信息
func DBGetOneRes( resId int)(resInfo ResourceTable, err error)  {
	err = db_conn.DB.Where("resource_id = ?",resId).First(&resInfo).Error;
	return
}

//根据resourceName获取一个资源的详细信息, 模糊查询
func DBGetResByName(resName string)(resInfo []ResourceTable, err error)  {
	err = db_conn.DB.Where("resource_name LIKE ?","%"+resName+"%").Find(&resInfo).Error;
	if err != nil {
		return nil, err
	}
	return
}

//根据tagId,获得同一类型资源
func DBGetAllResById(tagId int)(resInfo []ResourceTable, err error) {
	err = db_conn.DB.Where("tag_id = ?",tagId).Find(&resInfo).Error;
	if err != nil {
		return nil, err
	}
	return
}

//获取所有资源的信息
func DBGetAllRes()(resInfo []ResourceTable, err error) {
	err = db_conn.DB.Find(&resInfo).Error;
	if err != nil {
		return nil, err
	}
	return
}

//更新一条资源的信息
func DBUpdateRes(resInfo *ResourceTable)(err error){
	var res ResourceTable
	err = db_conn.DB.Model(&res).Where("resource_id = ?", resInfo.ResourceId).Updates(&resInfo).Error
	return
}
//创建一条新的资源
func DBCreateRes(resInfo *ResourceTable) (err error){
	err = db_conn.DB.Create(&resInfo).Error
	return
}
//删除一条资源
func DBDeleteRes(resId int) (err error){
	err = db_conn.DB.Where("resource_id = ?",resId).Delete(&ResourceTable{}).Error
	return
}