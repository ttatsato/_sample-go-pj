package domain

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title           string `json:"title" gorm:"not null" validate:"required"`
	Overview        string `json:"overview" sql:"type:text;" gorm:"not null" validate:"required"`
	Problem         string `json:"problem" gorm:"not null" validate:"required"`
	CustomerSegment string `json:"customerSegment" gorm:"not null" validate:"required"`
	Solution        string `json:"solution" gorm:"not null" validate:"required"`
	UniqueValue     string `json:"uniqueValue"`
	Channel         string `json:"channel"`
	KeyMetric       string `json:"keyMetric"`
	UnfairAdvantage string `json:"unfairAdvantage"`
	CostStructure   string `json:"costStructure"`
	RevenueStream   string `json:"revenueStream"`
	ElevatorPitch   string `json:"elevatorPitch" sql:"type:text;"`
	Remarks         string `json:"remarks" sql:"type:text;"`
	Status          string `json:"status"`
}
