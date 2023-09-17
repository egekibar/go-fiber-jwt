package models

type Users struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" valid:"required"`
	Email    string `json:"email" valid:"required,email" gorm:"unique"`
	Password string `json:"-" valid:"required"`
}
