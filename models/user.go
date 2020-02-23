package models

type Users struct {
	ID       int    `gorm:"primary_key:true" json:"id"`
	Name     string `gorm:"name" json:"name"`
	Email    string `gorm:"email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Phone    string `gorm:"phone" json:"phone"`
	Status   int    `gorm:"status" json:"status"`
}
