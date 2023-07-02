package models

import "gorm.io/gorm"

type WeekChapterType struct {
	gorm.Model
	BookID    uint
	ParshaID  uint8
	DayOfWeek uint8
	ContentID uint `gorm:"index"`
	Content   Content
}
