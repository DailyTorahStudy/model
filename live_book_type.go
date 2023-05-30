package models

import "gorm.io/gorm"

type LiveBookType struct {
	gorm.Model
	BookID    uint
	ChapterId int
	Name      string
	ContentID uint           `gorm:"index"`
	ParentID  *uint          `gorm:"index;autoIncrement:false"` // Обновите тип ParentID на *uint
	Parent    *LiveBookType  `gorm:"foreignkey:ParentID"`
	Children  []LiveBookType `gorm:"foreignkey:ParentID"`
	File      string
	Content   Content
}
