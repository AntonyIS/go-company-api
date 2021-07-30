package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" // import Gin Framework
)

// All request handlers and data are stored here

// Define the structure of a tech company
type TechFirm struct {
	Location string ` json "location" `
	Name     string ` json "name" `
	CEO      string ` json "ceo" `
	ID       int    ` json "id" `
}

// Define companies slice[dynamic array]
var companies = []TechFirm{
	{ID: 1, Location: "Menlo Park, USA", Name: "Facebook", CEO: "Mark Zuckerberg"},
	{ID: 2, Location: "Palo Alto, USA", Name: "Tesla", CEO: "Elon Musk"},
	{ID: 3, Location: "Seattle , USA", Name: "Amazon", CEO: "Andy Jassy"},
	{ID: 4, Location: "Redmond USA", Name: "MicroSoft", CEO: "Satya Nadella"},
	{ID: 5, Location: "Mountain View, USA", Name: "Google", CEO: "Sundra Pichai"},
	{ID: 6, Location: "Cupertino", Name: "Apple", CEO: "Tim Cook"},
}

func GetCompaniesHandler(c *gin.Context) {
	/*
	 Return slice of companies
	*/
	c.IndentedJSON(http.StatusOK, companies)
}

func GetCompanyHandler(c *gin.Context) {
	/*
		Returns company with the given ID else error
	*/
	// Get the ID of the company
	requestID := c.Param("id") // id type is string
	// Covert ID to int
	id, err := strconv.Atoi(requestID)

	// Check if conversion returns as error
	if err != nil {
		log.Fatalf("ERROR =====> %v", err)
	}
	// Loop through existing company slice and return company that matches the request id

	for _, company := range companies {
		if company.ID == id {
			c.IndentedJSON(http.StatusOK, company)
			return
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"statusCode": 404, "message": "Company not found"})

}

func PostCompanyHandler(c *gin.Context) {
	/*
		Receives request body of a new company and adds it to existing companies slice
	*/
	var newCompany TechFirm

	// Bind the received json to newCompany
	if err := c.BindJSON(&newCompany); err != nil {
		// Return an error, incoming json format does not match structure of TechFirm Struct
		log.Fatalf("ERROR =====> %v", err)
	}
	// Add new company to existing companies slice
	companies = append(companies, newCompany)

	c.IndentedJSON(http.StatusCreated, newCompany)
}

func EditCompanyHandler(c *gin.Context) {
	/*
		Returns edited company with the given ID else error
	*/
	var newCompany TechFirm

	// Bind the received json to newCompany
	if err := c.BindJSON(&newCompany); err != nil {
		// Return an error, incoming json format does not match structure of TechFirm Struct
		log.Fatalf("ERROR =====> %v", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}

	// Loop through companies
	for index, company := range companies {
		// Check if company exists with given from request
		// ID can not be changed
		fmt.Println(company.ID == newCompany.ID)
		if company.ID == newCompany.ID {
			// company to be edited found
			company.Location = newCompany.Location
			company.Name = newCompany.Name
			company.CEO = newCompany.CEO

			companies = append(companies[:index], companies[index+1:]...)
			// Add edited company to existing companies slice
			companies = append(companies, newCompany)
			c.IndentedJSON(http.StatusCreated, newCompany)
			return
		}

	}
}

func DeleteCompanyHandler(c *gin.Context) {
	/*
		Deletes company from slice
	*/
	// Get the ID of the company
	requestID := c.Param("id") // id type is string
	// Covert ID to int
	id, err := strconv.Atoi(requestID)

	// Check if conversion returns as error
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Internal server error"})
	}
	// Loop through existing company slice and return company that matches the request id

	for index, company := range companies {
		if company.ID == id {
			companies = append(companies[:index], companies[index+1:]...)
		}
	}
	c.IndentedJSON(http.StatusOK, companies)
}

var Router = gin.Default()
