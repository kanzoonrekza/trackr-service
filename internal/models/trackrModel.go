package models

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Trackr struct {
	gorm.Model
	UserID uint `gorm:"not null"`

	Src       string `gorm:"not null"`
	SrcId     string `gorm:"not null"`
	SrcStatus string `gorm:"not null"`

	Title          string         `gorm:"not null"`
	AltTitle       datatypes.JSON `gorm:"type:jsonb"`
	TotalEpisode   uint16         `gorm:"not null; default:1"`
	CurrentEpisode uint16         `gorm:"not null; default:1"`
	Status         string         `gorm:"not null; default:'Plan to Watch'"`
	IsCompleted    bool           `gorm:"not null; default:false"`
	PersonalScore  int8           `gorm:"not null"`

	Category string
	Tags     pq.StringArray `gorm:"type:text[]"`
}
