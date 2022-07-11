package main

import (
	"ngouzo/ngouzo-backend/infrastructure"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	router := infrastructure.NgRouter()
	// db := infrastructure.NgDataBase()

	// if err := db.DB.AutoMigrate(); err != nil {
	// 	errors.New("Unable autoMigrateDB - " + err.Error())
	// }

	router.Gin.Run(":7070")
}
