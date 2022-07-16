package models

import "time"

type StudentMarksEvaluation struct {
	ID                 int64     `gorm:"primary_key;auto_increment"  json:"id"`
	StudentID          int       `json:"student_id"`
	Students           Students  `gorm:"foreignKey:StudentID"`
	TeacherID          int       `json:"teacher_id"`
	Teacher            User      `gorm:"foreignKey:User"`
	ModuleID           int       `json:"module_id"`
	Modules            Modules   `gorm:"foreignKey:ModuleID"`
	StudentScore       float64   `json:"student_score"`
	StudentGrade       float64   `json:"student_grades"`
	StuddingRangeLevel string    `json:"studding_range_year"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}

func (studentMarksEvaluation *StudentMarksEvaluation) TableName() string {
	return "StudentMarksEvaluation"
}

type RecordStudentMarksEvaluation struct {
	StudentID    int     `form:"student_id" binding:"required"`
	StudentScore float64 `form:"student_score" binding:"required"`
	TeacherID    int     `form:"teacher_id" binding:"required"`
	StudentGrade float64 `form:"student_grades"`
	ModuleID     int     `form:"module_id"`
}

type UpdateStudentMarksEvaluation struct {
	ID           int64   `form:"id" binding:"required"`
	StudentScore float64 `form:"student_score"`
	StudentGrade float64 `form:"student_grades"`
	ModuleID     int     `form:"module_id"`
}

func (studentMarksEvaluation *StudentMarksEvaluation) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = studentMarksEvaluation.ID
	resp["student_score"] = studentMarksEvaluation.StudentScore
	resp["student_grades"] = studentMarksEvaluation.StudentGrade
	resp["teacher_id"] = studentMarksEvaluation.TeacherID
	resp["student_id"] = studentMarksEvaluation.StudentID
	resp["module_id"] = studentMarksEvaluation.ModuleID

	return resp
}
