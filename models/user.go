package models

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"size:100;unique;not null" json:"username"`
	Email    string `gorm:"size:100;unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"` // stored in DB, hidden in API responses
}