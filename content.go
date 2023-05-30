package models

import (
	"errors"
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	ParentID         *int
	Order            uint8
	ContentParagraph []ContentParagraph
}

type ContentParagraph struct {
	gorm.Model
	ContentID uint
	Order     uint16
	Lang      string
	Html      string
}

// BeforeCreate является пользовательским методом GORM, который вызывается перед сохранением записи в базе данных.
func (c *ContentParagraph) BeforeCreate(tx *gorm.DB) (err error) {
	if c.Lang != "ru" && c.Lang != "he" {
		// Если значение Lang не "ru" или "he", отклоняем сохранение записи и возвращаем ошибку
		err = errors.New("invalid value for Lang, allowed values are 'ru' and 'he'")
	}
	return
}
