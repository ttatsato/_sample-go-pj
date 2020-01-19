package domain

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title   string  `json:"title" gorm:"not null"`
	Overview string  `json:"overview" sql:"type:text;" gorm:"not null"`
	Problem string  `json:"problem" gorm:"not null"`
	CustomerSegment  string  `json:"customerSegment" gorm:"not null"`
	Solution  string  `json:"solution" gorm:"not null"`
	UniqueValue string  `json:"uniqueValue"`
	Channel  string  `json:"channel"`
	KeyMetric  string  `json:"keyMetric"`
	UnfairAdvantage  string  `json:"unfairAdvantage"`
	CostStructure  string  `json:"costStructure"`
	RevenueStream  string  `json:"revenueStream"`
	ElevatorPitch string  `json:"elevatorPitch" sql:"type:text;"`
	Remarks string `json:"remarks" sql:"type:text;"`
	Status string `json:"status"`
}