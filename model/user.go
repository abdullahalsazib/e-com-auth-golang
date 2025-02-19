package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"column:name"`
	Email     string    `json:"email" gorm:"column:email;unique"`
	Password  string    `json:"-" gorm:"column:password;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}
