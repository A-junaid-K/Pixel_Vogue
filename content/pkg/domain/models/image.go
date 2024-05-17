package models

import "time"

type Image struct {
	Id      int          `json:"id" gorm:"primaryKey;autoIncrement"`
	Image   string       `json:"image,omitempty"`
	Details ImageDetails `json:"imagedetails,omitempty" gorm:"foreignKey:ImageID"`
}

type ImageDetails struct {
	Id              int    `json:"id" gorm:"primaryKey;autoIncrement"`
	ContributorId   int    `json:"contributor_id"`
	Size            string `json:"size,omitempty"`
	Dimension      string
	DateTaken       time.Time
	MoreInformation string
	Tags            string
	ImageID         int `json:"image_id"`
}

type ImageResonse struct {
	Like    int
	Comment string
}
