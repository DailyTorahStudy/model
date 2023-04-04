package models

import "gorm.io/gorm"

type TehelimDayType struct {
	gorm.Model
	BookID    uint
	Day       uint8
	Order     uint8
	ContentID uint `gorm:"index"`
	Content   Content
}

type TehelimNumberType struct {
	gorm.Model
	BookID    uint
	Number    uint8
	ContentID uint `gorm:"index"`
	Content   Content
}
