/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/17 12:42 上午
 * @Description:
 */

package tag

import (
	"folk_culture_api/db_conn"
)

type TagsTable struct {
	TagId          int
	TagName        string
	TagSize        int
	TagContext     string
}

// 设置TagsTable 的表名为`tags_table`                            //⚠️非常值得注意：数据库表名与model的映射关系
func (TagsTable) TableName() string {
	return "tags_table"
}


//根据tagid获取一个分类标签详细信息
func DBGetOneTag(tagId int)(userInfo TagsTable, err error)  {
	err = db_conn.DB.Where("tag_id = ?",tagId).First(&userInfo).Error;
	return
}

//获取所有分类标签的信息
func DBGetAllTag()(tagsInfo []TagsTable, err error) {
	err = db_conn.DB.First(&tagsInfo).Error;
	if err != nil {
		return nil, err
	}
	return
}

//更新分类标签的数量信息
func DBUpdateTag(tagInfo *TagsTable)(err error){
	var tag TagsTable
	err = db_conn.DB.Model(&tag).Where("tag_id = ?", tagInfo.TagId).Updates(&tagInfo).Error
	return
}

func DBCreateTag(tagInfo *TagsTable) (err error){
	err = db_conn.DB.Create(&tagInfo).Error
	return
}

func DBDeleteTag(tagId int) (err error){
	err = db_conn.DB.Where("tag_id = ?",tagId).Delete(&TagsTable{}).Error
	return
}