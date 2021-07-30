package main

import (
	"example.com/routes"
)

func main() {
	// The API starts running from the main function
	router := routes.Router

	// Get company slice

	router.GET("/companies", routes.GetCompaniesHandler)
	router.GET("/companies/:id", routes.GetCompanyHandler)
	router.POST("/companies", routes.PostCompanyHandler)
	router.PUT("companies/:id", routes.EditCompanyHandler)
	router.DELETE("/companies/:id", routes.DeleteCompanyHandler)

	// Run server
	router.Run("localhost:8080")
}
