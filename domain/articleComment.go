package domain

import "github.com/jinzhu/gorm"

type ArticleComment struct {
	gorm.Model
	ArticleId uint   `json:"articleId" gorm:"not null"`
	Text      string `json:"text" sql:"type:text;" gorm:"not null" validate:"required"`
	Status    string `json:"status"`
}
