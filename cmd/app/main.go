package main

import (
	_ "github.com/BeRebornBng/OsauAmsApi/docs"
	"github.com/BeRebornBng/OsauAmsApi/internal/app"
)

// @title Osau AMS API
// @version 1.0
// @description This is the API for Osau AMS.

// @contact.name API Support
// @contact.email bakit.german.work@icloud.com

// @host localhost:8000
// @BasePath /api

// @tag.name Users
// @tag.description Endpoints for Admin role

// @tag.name Headman
// @tag.description Endpoints for Headman role

// @tag.name Student
// @tag.description Endpoints for Student role

// @tag.name Teacher
// @tag.description Endpoints for Teacher role

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app.Run()
}
