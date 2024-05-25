package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	CartID  int `json:"cart_id" gorm:"primaryKey;autoIncrement"`
	UserID  int `json:"user_id,omitempty"`
	ImageID int `json:"image_id,omitempty"`
}

type CartItem struct {
	Id         uint
	CartID     uint
	ImageID    uint
	Price      float64
	Created_At time.Time
	Updated_At time.Time
	Deleted_At gorm.DeletedAt
}
