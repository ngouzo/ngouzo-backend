package routes

import (
	"ngouzo/ngouzo-backend/api/controllers"
	"ngouzo/ngouzo-backend/infrastructure"
)

type ClassesRoute struct {
	Handler    infrastructure.GinRouter
	Controller controllers.ClassesController
}

func ClassRoute(
	controller controllers.ClassesController,
	handler infrastructure.GinRouter,
) ClassesRoute {
	return ClassesRoute{
		Handler:    handler,
		Controller: controller,
	}
}

func (cl ClassesRoute) Setup() {
	class := cl.Handler.Gin.Group("/class")
	{
		class.POST("/record-new-class", cl.Controller.RecordClasses)
		class.PUT("/update-class-name", cl.Controller.UpdateClasses)
		class.GET("/retrieve-classes", cl.Controller.GetClasses)
	}
}
