package main

import (
	"github.com/lutfiharidha/haioo/app"
	"github.com/lutfiharidha/haioo/database"
)

// @title           Haioo Cart Service
// @version         1.0
// @description     Haioo Cart Service
// @host      		localhost:8081
// @in                          header
func main() {
	db := app.NewSQL().SetupDatabaseConnection() //setup database connection

	database.Migrator(db) //migrate table
	app.Router()          //start server
}
