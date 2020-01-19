package domain

import "github.com/jinzhu/gorm"

type ArticleComment struct {
	gorm.Model
	ArticleID uint   `json:"articleId"`
	Text      string `json:"text" sql:"type:text;" gorm:"not null" validate:"required"`
}
