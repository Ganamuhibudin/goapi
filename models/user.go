package models

type Users struct {
	ID     int    `gorm:"primary_key:true" json:"id"`
	Name   string `gorm:"name" json:"name"`
	Email  string `gorm:"email" json:"email"`
	Phone  string `gorm:"phone" json:"subject_template"`
	Status int    `gorm:"status" json:"status"`
}
