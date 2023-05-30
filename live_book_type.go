package models

import "gorm.io/gorm"

type LiveBookType struct {
	gorm.Model
	BookID    uint
	ChapterId int
	Name      string
	ContentID uint           `gorm:"index"`
	ParentID  uint           `gorm:"index"` // Добавлено поле ParentID
	Parent    *LiveBookType  `gorm:"foreignkey:ParentID"`
	Children  []LiveBookType `gorm:"foreignkey:ParentID"`
	Content   Content
}
