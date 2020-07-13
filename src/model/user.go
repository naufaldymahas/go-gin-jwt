package model

import "time"

type User struct {
	ID        int64     `gorm:"PRIMARY_KEY;NOT NULL;UNIQUE" json:"id"`
	Name      string    `gorm:"NOT NULL" json:"name"`
	Email     string    `gorm:"NOT NULL;UNIQUE" json:"email"`
	Password  string    `gorm:"NOT NULL" json:"password"`
	CreatedAt time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP" json:"updated_at"`
}
