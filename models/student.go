package models

import "time"

type Students struct {
	ID                 int64     `gorm:"primary_key;auto_increment" json:"id"`
	SchoolID           int       `json:"School_id"`
	School             User      `gorm:"foreignKey:SchoolID"`
	FirstName          string    `gorm:"first_name"`
	LastName           string    `gorm:"last_name"`
	Picture            string    `gorm:"picture"`
	BirthDate          string    `gorm:"birth_date"`
	Gender             string    `gorm:"gender"`
	BirthCountry       string    `gorm:"birth_country"`
	BirthCityOrVillage string    `gorm:"birth_city_or_village"`
	IsCurrentStudent   bool      `gorm:"is_current_student"`
	LeftAt             time.Time `json:"left_at,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdateAt           time.Time `json:"updated_at,omitempty"`
}

func (student *Students) TableName() string {
	return "Students"
}

type RegisterStudents struct {
	SchoolID           int    `form:"School_id"`
	FirstName          string `form:"first_name"`
	LastName           string `form:"last_name"`
	Picture            string `form:"picture"`
	Gender             string `gorm:"gender"`
	BirthCountry       string `gorm:"birth_country"`
	BirthCityOrVillage string `gorm:"birth_city_or_village"`
}

type UpdateStudents struct {
	ID                 int64  `form:"id" binding:"required"`
	FirstName          string `form:"first_name"`
	LastName           string `form:"last_name"`
	Picture            string `form:"picture"`
	Gender             string `gorm:"gender"`
	BirthCountry       string `gorm:"birth_country"`
	BirthCityOrVillage string `gorm:"birth_city_or_village"`
}

func (student *Students) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = student.ID
	resp["school_id"] = student.SchoolID
	resp["first_name"] = student.FirstName
	resp["last_name"] = student.LastName
	resp["picture"] = student.Picture
	resp["gender"] = student.Gender
	resp["birth_country"] = student.BirthCountry
	resp["birth_city_or_village"] = student.BirthCityOrVillage

	return resp
}
