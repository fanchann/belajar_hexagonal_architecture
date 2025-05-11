package entity

type Books struct {
	ID          string `gorm:"primaryKey;column:id"`
	Title       string `gorm:"column:title"`
	Author      string `gorm:"column:author"`
	Description string `gorm:"column:description"`
}

func (b *Books) TableName() string {
	return "books"
}
