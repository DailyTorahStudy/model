package models

import "gorm.io/gorm"

type LiveBookType struct {
	gorm.Model
	BookID    uint
	ContentID uint `gorm:"index"`
	Content   Content
}
