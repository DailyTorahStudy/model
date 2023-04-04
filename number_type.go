package models

import "gorm.io/gorm"

type NumberType struct {
	gorm.Model
	BookID    uint
	Number    uint16
	ContentID uint `gorm:"index"`
	Content   Content
}
