package models

import "gorm.io/gorm"

type WeekChapterType struct {
	gorm.Model
	BookID    uint  `gorm:"index"`
	ParshaID  uint8 `gorm:"index"`
	DayOfWeek uint8 `gorm:"index"`
	ContentID uint  `gorm:"index"`
	Content   Content
}
