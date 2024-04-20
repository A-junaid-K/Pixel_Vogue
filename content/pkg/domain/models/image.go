package models

import "time"

type Image struct {
	Id      int          `json:"id" gorm:"primaryKey;autoIncrement"`
	Image   string       `json:"image,omitempty"`
	Details ImageDetails `json:"imagedetails,omitempty" gorm:"foreignKey:ImageID"`
}

type ImageDetails struct {
	// Contributor      Contributor
	Id               int    `json:"id" gorm:"primaryKey;autoIncrement"`
	ImageID          int    `json:"image_id"`
	Size             string `json:"size,omitempty"`
	Dimensions       string
	Date_Taken       time.Time
	More_Information string
	Tagss            string
}

type ImageResonse struct {
	Like    int
	Comment string
}
