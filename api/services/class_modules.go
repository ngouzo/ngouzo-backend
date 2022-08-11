package services

import (
	"ngouzo/ngouzo-backend/api/repositories"
	"ngouzo/ngouzo-backend/models"
)

type ClassModuleService struct {
	repo repositories.ClassModuleRepository
}

func RecordClassService(repo repositories.ClassModuleRepository) ClassModuleService {
	return ClassModuleService{
		repo: repo,
	}
}

func (c ClassModuleService) RecordClassModules(classM models.RecordClassModules) error {
	return c.repo.RecordClassModule(classM)
}

func (c ClassModuleService) UpdateClassModules(classM models.UpdateClassModules) error {
	return c.repo.UpdateClassModules(classM)
}

func (c ClassModuleService) GetClassModule() error {
	return c.repo.GetClassModule()
}
