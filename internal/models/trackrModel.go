package models

import "gorm.io/gorm"

type Trackr struct {
	gorm.Model
	Title          string `gorm:"not null"`
	TotalEpisode   uint16 `gorm:"not null"`
	CurrentEpisode uint16 `gorm:"not null; default:1"`
	Status         string `gorm:"not null; default:'Plan to Watch'"`
	Completed      bool   `gorm:"not null; default:false"`
	Rate           int8   `gorm:"not null; check:Rate >= 0 AND Rate <= 10"`
	UserID         uint   `gorm:"not null"`
}
