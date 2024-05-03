package models

import "time"

type Image struct {
	Id      int          `json:"id" gorm:"primaryKey;autoIncrement"`
	Image   string       `json:"image,omitempty"`
	Details ImageDetails `json:"imagedetails,omitempty" gorm:"foreignKey:ImageID"`
}

type ImageDetails struct {
	ContributorId   int    `json:"contributor_id"`
	Id              int    `json:"id" gorm:"primaryKey;autoIncrement"`
	ImageID         int    `json:"image_id"`
	Size            string `json:"size,omitempty"`
	Dimensions      string
	DateTaken       time.Time
	MoreInformation string
	Tags            string
}

type ImageResonse struct {
	Like    int
	Comment string
}
