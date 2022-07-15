package models

import "time"

type User struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"name"`
	FirstName   string    `gorm:"first_name"`
	LastName    string    `gorm:"last_name"`
	Email       string    `gorm:"email"`
	PhoneNumber string    `gorm:"phone_number"`
	Logo        string    `gorm:"logo"`
	Username    string    `gorm:"username"`
	Password    string    `gorm:"password"`
	IsAdmin     bool      `gorm:"is_school"`
	IsUser      bool      `gorm:"is_user"`
	IsActive    bool      `gorm:"is_active"`
	Confirmed   bool      `gorm:"confirmed"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdateAt    time.Time `json:"updated_at,omitempty"`
}

func (user *User) TableName() string {
	return "User"
}

type RegisterUser struct {
	Name        string `form:"name"`
	FirstName   string `form:"first_name"`
	LastName    string `form:"last_name"`
	Email       string `form:"email" binding:"required"`
	PhoneNumber string `form:"phone_number"`
	Logo        string `form:"logo"`
	Username    string `form:"username"  binding:"required"`
	Password    string `form:"password" binding:"required"`
}

type UpdateUser struct {
	ID          int64  `form:"id" binding:"required"`
	Name        string `form:"name"`
	FirstName   string `form:"first_name"`
	LastName    string `form:"last_name"`
	Email       string `form:"email" binding:"required"`
	PhoneNumber string `form:"phone_number"`
	Logo        string `form:"logo"`
	Username    string `form:"username"  binding:"required"`
	Password    string `form:"password" binding:"required"`
}

func (user *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = user.ID
	resp["name"] = user.Name
	resp["first_name"] = user.FirstName
	resp["last_name"] = user.LastName
	resp["email"] = user.Email
	resp["phone_number"] = user.PhoneNumber
	resp["logo"] = user.Logo
	resp["username"] = user.Username
	resp["password"] = user.Password
	resp["is_admin"] = user.IsAdmin
	resp["is_user"] = user.IsUser
	resp["is_active"] = user.IsActive
	resp["confirmed"] = user.Confirmed

	return resp
}
