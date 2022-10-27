package user

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;column:id"`
	Username  string    `gorm:"not null;column:username;uniqueIndex"`
	Email     string    `gorm:"not null;column:email;uniqueIndex"`
	Password  string    `gorm:"not null;column:password"`
	Age       uint      `gorm:"not null;column:age"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
