package main

import (
	"errors"
	"ngouzo/ngouzo-backend/api/controllers"
	"ngouzo/ngouzo-backend/api/repositories"
	"ngouzo/ngouzo-backend/api/routes"
	"ngouzo/ngouzo-backend/api/services"
	"ngouzo/ngouzo-backend/infrastructure"
	"ngouzo/ngouzo-backend/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	router := infrastructure.NgRouter()
	db := infrastructure.NgDataBase()

	// Classes
	classRepository := repositories.RecordClassesRepository(db)
	classService := services.RecordClassesService(classRepository)
	classController := controllers.RecordClassesController(classService)
	classRoute := routes.ClassRoute(classController, router)
	classRoute.Setup()

	if err := db.DB.AutoMigrate(&models.Classes{}); err != nil {
		errors.New("Unable autoMigrateDB - " + err.Error())
	}

	router.Gin.Run(":7070")
}
