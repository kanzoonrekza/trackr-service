package models

import "gorm.io/gorm"

type Trackr struct {
	gorm.Model
	Title          string
	TotalEpisode   int
	CurrentEpisode int
}
