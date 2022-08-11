package services

import (
	"ngouzo/ngouzo-backend/api/repositories"
	"ngouzo/ngouzo-backend/models"
)

type ClassesService struct {
	repo repositories.ClassesRepository
}

func RecordClassesService(repo repositories.ClassesRepository) ClassesService {
	return ClassesService{
		repo: repo,
	}
}

func (c ClassesService) RecordClasses(class models.RecordClasses) error {
	return c.repo.RecordClasses(class)
}

func (c ClassesService) UpdateClasses(class models.UpdateClasses) error {
	return c.repo.UpdateClasses(class)
}

func (c ClassesService) GetClasses() error {
	return c.repo.GetClasses()
}
