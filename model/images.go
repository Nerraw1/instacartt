package model

type Images struct {
	Image_id   uint   `json:"image_id" gorm:"primaryKey"`
	Image_name string ` json:"imag_name"`
	Image_url  string `json:"image_url"`
}
