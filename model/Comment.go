package model

import "gorm.io/gorm"

type Comment struct{
	gorm.Model
	UserID	uint	`json:"user_id"`
	ArticleID string `gorm:"type:int unsigned;not null" json:"article_id"`	
	RandomName string	`gorm:"type:varchar(20);nut null" json:"random_name"`
	Content string	`gorm:"type:varchar(500);not null" json:"content"`
}