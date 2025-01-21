package model

type User struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"type:text;not null"`
	Email string `gorm:"type:text;not null"`
}
