package models

import "gorm.io/gorm"

// Book Тут перечисляются все книги которые мы показываем, например ЙОМ-ЙОМ, ХУМАШ и т.д.
type Book struct {
	gorm.Model
	Code              string
	BookName          string
	BookTypeID        string `gorm:"foreignKey:ID"`
	WeekChaptersTypes []WeekChapterType
	CalendarTypes     []YearType
	NumberTypes       []NumberType
	LiveBookTypes     []LiveBookType
}
