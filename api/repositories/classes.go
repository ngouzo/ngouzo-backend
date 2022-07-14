package repositories

import (
	"ngouzo/ngouzo-backend/infrastructure"
	"ngouzo/ngouzo-backend/models"
)

type ClassesRepository struct {
	db infrastructure.Database
}

func RecordClassesRepository(db infrastructure.Database) ClassesRepository {
	return ClassesRepository{
		db: db,
	}
}

func (c ClassesRepository) RecordClasses(class models.RecordClasses) error {
	var dbClasses models.Classes
	dbClasses.Name = class.Name

	err := c.db.DB.Where(map[string]interface{}{"name": class.Name}).Find(&dbClasses).Error

	if err != nil {
		return err
	}

	return c.db.DB.Create(&dbClasses).Error
}

func (c ClassesRepository) UpdateClasses(class models.UpdateClasses) error {
	var dbClasses models.Classes
	dbClasses.Name = class.Name

	err := c.db.DB.Model(&dbClasses).Where("id = ?", class.ID).Error

	if err != nil {
		return err
	}

	c.db.DB.Model(&dbClasses).Where("id = ?", class.ID).Update("name", class.Name)
	return nil
}

func (c ClassesRepository) GetClasses() error {
	var dbClasses models.Classes

	err := c.db.DB.Find(&dbClasses).Error
	return err
}
