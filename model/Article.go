package model

import (
	"fmt"
	"myapp/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	User         User   `gorm:"foreignKey:UserID"`                         //关联用户
	UserID       int    `gorm:"type:int unsigned;not null" json:"user_id"` //用户id
	Title        string `gorm:"type:varchar(100);not null" json:"title"`   //标题
	Content      string `gorm:"type:longtext" json:"content"`              //帖子主体
	CommentCount int    `gorm:"type:int" json:"comment_count"`             //评论数量
	Like         int    `gorm:"type:int" json:"like"`                      //点赞
	UnLike       int    `gorm:"type:int" json:"unlike"`                    //踩
	// Img        string   `gorm:"type:varchar(100)" json:"img"`  //帖子图片
}

// 创建帖子
func CreateArt(data *Article) int {
	err := Db.Create(&data).Error
	fmt.Println(data)
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

// GetArt 查询帖子列表
func GetArt(pageSize int, pageNum int) ([]Article, int) {
	var artList []Article
	err = Db.Preload("User").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&artList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return artList, errmsg.SUCCESS
}

//编辑 todo

// 删除帖子
func DeleteArt(id int) int {
	var art Article

	err := Db.Where("id=?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
