package models

import "gorm.io/gorm"

type LiveBookType struct {
	gorm.Model
	BookID    uint
	ChapterId int
	Name      string
	ContentID uint           `gorm:"index"`
	Parent    *LiveBookType  `gorm:"foreignkey:ParentID"`
	Children  []LiveBookType `gorm:"foreignkey:ParentID"`
	Content   Content
}
