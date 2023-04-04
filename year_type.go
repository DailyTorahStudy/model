package models

import "gorm.io/gorm"

type YearType struct {
	gorm.Model
	BookID    uint
	Day       uint8
	Month     uint8
	Year      uint16
	IsLeap    bool
	ContentID uint `gorm:"index"`
	Content   Content
}
