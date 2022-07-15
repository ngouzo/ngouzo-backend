package models

import "time"

type UserProfile struct {
	ID           int64     `gorm:"primary_key;auto_increment" json:"id"`
	UserID       int       `json:"user_id"`
	User         User      `gorm:"foreignKey:UserID"`
	Sex          string    `gorm:"sex"`
	Picture      string    `gorm:"picture"`
	BirthDate    string    `gorm:"birth_date"`
	IsTeacher    bool      `gorm:"is_teacher"`
	IsAccounting bool      `gorm:"is_accounting"`
	IsDriver     bool      `gorm:"is_driver"`
	IsCashier    bool      `gorm:"is_cashier"`
	IsParent     bool      `gorm:"is_parent"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

func (userProfile *UserProfile) TableName() string {
	return "UserProfile"
}

type SaveUserProfile struct {
	Sex       string `form:"sex"`
	BirthDate string `form:"birth_date"`
	Picture   string `form:"picture"`
}

type UpdateUserProfile struct {
	UserID    int    `form:"user_id" binding:"required"`
	Sex       string `form:"sex"`
	BirthDate string `form:"birth_date"`
	Picture   string `form:"picture"`
}

func (userProfile *UserProfile) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = userProfile.ID
	resp["user_id"] = userProfile.UserID
	resp["sex"] = userProfile.Picture
	resp["birth_date"] = userProfile.BirthDate

	return resp
}
