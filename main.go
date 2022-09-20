package main

import (
	handler "github.com/AntonyIS/GO-REST-API-1/handler"
)

func main() {
	router := gin.Default()

	router.GET("/api/v1/companies", handler.GetCompaniesHandler)
	router.GET("/api/v1/companies/:id", handler.GetCompanyHandler)
	router.POST("/api/v1/companies", handler.PostCompanyHandler)
	router.PUT("/api/v1/companies/:id", handler.EditCompanyHandler)
	router.DELETE("/api/v1/companies/:id", handler.DeleteCompanyHandler)

	// Run server
	router.Run("localhost:8080")
}
