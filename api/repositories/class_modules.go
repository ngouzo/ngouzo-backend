package repositories

import (
	"ngouzo/ngouzo-backend/infrastructure"
	"ngouzo/ngouzo-backend/models"
)

type ClassModuleRepository struct {
	db infrastructure.Database
}

func RecordClassModuleRepository(db infrastructure.Database) ClassModuleRepository {
	return ClassModuleRepository{
		db: db,
	}
}

func (c ClassModuleRepository) RecordClassModule(classM models.RecordClassModules) error {
	var dbClassModule models.ClassModules
	dbClassModule.Grades = classM.Grades
	dbClassModule.Hours = classM.Hours
	dbClassModule.ClassID = classM.ClassID
	dbClassModule.ModuleID = classM.ModuleID
	dbClassModule.SchoolID = classM.SchoolID

	err := c.db.DB.Where(map[string]interface{}{"class_id": classM.ClassID, "module_id": classM.ModuleID, "school_id": classM.SchoolID}).Find(&dbClassModule).Error

	if err != nil {
		return err
	}

	return c.db.DB.Create(&dbClassModule).Error
}

func (c ClassModuleRepository) UpdateClassModules(classM models.UpdateClassModules) error {
	var dbClassModule models.ClassModules
	dbClassModule.Grades = classM.Grades
	dbClassModule.Hours = classM.Hours
	dbClassModule.ClassID = classM.ClassID
	dbClassModule.ModuleID = classM.ModuleID

	err := c.db.DB.Model(&dbClassModule).Where("id = ? ", classM.ID).Error

	if err != nil {
		return err
	}

	c.db.DB.Model(&dbClassModule).Where("id = ?", classM.ID).Updates(
		map[string]interface{}{
			"class_id":  classM.ClassID,
			"module_id": classM.ModuleID,
			"grades":    classM.Grades,
			"hours":     classM.Hours})
	return nil

}

func (c ClassModuleRepository) GetClassModule() error {
	var dbClassModule models.ClassModules

	err := c.db.DB.Find(&dbClassModule).Error
	return err
}
