package models

import "time"

type User struct {
	Id         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Email      string `json:"email,omitempty" validate:"email"`
	Password   string `json:"password,omitempty" validate:"min=6"`
	Is_blocked bool   `json:"isblocked" gorm:"default=false"`
	Role       string `json:"role" gorm:"NOT NULL"`
	Validate   bool   `json:"validate" gorm:"NOT NULL; default:false"`
	Created_at time.Time
	Otp        int
}

type UserProfile struct{
	Bio string `json:"bio"`

}