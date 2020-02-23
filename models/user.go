package models

import "time"

type Users struct {
	ID            int        `gorm:"primary_key:true" json:"id"`
	Name          string     `gorm:"name" json:"name"`
	Email         string     `gorm:"email" json:"email"`
	Password      string     `gorm:"column:password" json:"password"`
	VerifiedEmail int8       `gorm:"column:verified_email" json:"verified_email"`
	Phone         string     `gorm:"phone" json:"phone"`
	Status        int        `gorm:"status" json:"status"`
	LastLogin     *time.Time `gorm:"column:last_login_at" json:"last_login"`
	CreatedAt     time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"DEFAULT:NULL"`
}
