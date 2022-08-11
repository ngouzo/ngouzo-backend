package controllers

import (
	"net/http"
	"ngouzo/ngouzo-backend/api/services"
	"ngouzo/ngouzo-backend/models"
	"ngouzo/ngouzo-backend/utils"

	"github.com/gin-gonic/gin"
)

type ClassModuleController struct {
	service services.ClassModuleService
}

func RecordClassModuleController(s services.ClassModuleService) ClassModuleController {
	return ClassModuleController{
		service: s,
	}
}

func (cl *ClassModuleController) RecordClassesModule(c *gin.Context) {
	var classM models.RecordClassModules

	if err := c.ShouldBind(&classM); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
	}

	cl.service.RecordClassModules(classM)
	utils.SuccessJSON(c, http.StatusOK, "Module & Grades added with success")

}

func (cl *ClassModuleController) UpdateClassModules(c *gin.Context) {
	var classM models.UpdateClassModules

	if err := c.ShouldBind(&classM); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
	}

	err := cl.service.UpdateClassModules(classM)

	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Error updating data entered")
		return
	}

	cl.service.UpdateClassModules(classM)
	utils.SuccessJSON(c, http.StatusOK, "Class info updated with success")
}
