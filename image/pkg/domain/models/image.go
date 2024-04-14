package models

import "time"


type Image struct{
	Contributor Contributor
	Image string `json:"image,omitempty"`
	Size string `json:"size,omitempty"`
	Dimensions string
	Date_Taken time.Time
	More_Information string
	Tagss string
}

type ImageResonse struct{
	Like int
	Comment string
}