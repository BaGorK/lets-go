// Package main implements the main entry point for the REST API application.
package main

import (
	"github.com/BaGorK/lets-go/rest-api/db"
	"github.com/BaGorK/lets-go/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	db.InitDB()

	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run()
}
