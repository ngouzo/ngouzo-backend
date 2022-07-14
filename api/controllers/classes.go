package controllers

import (
	"net/http"
	"ngouzo/ngouzo-backend/api/services"
	"ngouzo/ngouzo-backend/models"

	"ngouzo/ngouzo-backend/utils"

	"github.com/gin-gonic/gin"
)

type ClassesController struct {
	service services.ClassesService
}

func RecordClassesController(s services.ClassesService) ClassesController {
	return ClassesController{
		service: s,
	}
}

func (cl *ClassesController) RecordClasses(c *gin.Context) {
	var class models.RecordClasses

	if err := c.ShouldBind(&class); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
	}

	cl.service.RecordClasses(class)
	utils.SuccessJSON(c, http.StatusOK, "Class name saved with success")
}

func (cl *ClassesController) UpdateClasses(c *gin.Context) {
	var class models.UpdateClasses

	if err := c.ShouldBind(&class); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
	}

	err := cl.service.UpdateClasses(class)

	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Error updating class name")
		return
	}

	cl.service.UpdateClasses(class)
	utils.SuccessJSON(c, http.StatusOK, "Class name updated with success")
}

func (cl *ClassesController) GetClasses(c *gin.Context) {
	var dbClasses models.Classes

	if err := c.ShouldBindJSON(&dbClasses); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
		return
	}

	c.JSON(http.StatusOK, cl.service.GetClasses())

}
