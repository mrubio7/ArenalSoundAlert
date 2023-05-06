package models

import (
	"time"
)

type MediaData struct {
	ID         int    `json:"id" gorm:"id"`
	DateString string `json:"date" gorm:"date"`
	Date       time.Time
	GUID       struct {
		Rendered string `json:"rendered" gorm:"image"`
	} `json:"guid"`
}

type Media struct {
	ID    int       `json:"id" gorm:"id"`
	Date  time.Time `json:"date" gorm:"date"`
	Image string    `json:"image" gorm:"image"`
}
