package models

// BookType Тут перечислены типы книг: "Year", "WeekChapters", "Number", "Book"
type BookType struct {
	ID    string `gorm:"type:varchar(50);primaryKey"`
	Books []Book
}
