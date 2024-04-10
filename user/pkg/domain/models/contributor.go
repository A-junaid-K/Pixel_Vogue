package models

import "time"

type Contributor struct {
	Id         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	FullName       string `json:"name,omitempty" validate:"min=3,max=20"`
	Email      string `json:"email,omitempty" validate:"email"`
	Password   string `json:"password,omitempty" validate:"min=6"`
	Is_blocked bool   `json:"isblocked" gorm:"default=false"`
	Role       string `json:"role" gorm:"NOT NULL"`
	Validate   bool   `json:"validate" gorm:"NOT NULL; default:false"`
	Created_at time.Time
	Otp        int
}

type ContributorProfile struct{
	Bio string `json:"bio"`
	Country string `json:"country,omitempty"`
	Street string `json:"street,omitempty"`
	City string `json:"city,omitempty"`
	State string `json:"state"`
	ZipCode int `json:"zip_code.omitempty" validate:"min=6,max=6"`
}

type ContributorUpdate struct{
	ProfilePhoto string
	CoverPhoto string
	Name string
	Username string 
	Bio string
	Email string
	Password string
}

