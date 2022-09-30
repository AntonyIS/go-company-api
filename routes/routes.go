package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Company struct {
	Location string `json:"location"`
	Name     string `json:"name"`
	CEO      string `json:"ceo"`
	ID       string `json:"id"`
}

func Router() {
	var router = gin.Default()

	router.GET("/", Home)
	router.GET("/companies", GetCompanies)
	router.GET("/companies/:id", GetCompany)
	router.POST("/companies", PostCompany)
	router.PUT("/companies/:id", EditCompany)
	router.DELETE("/companies/:id", DeleteCompany)
	// Run server
	router.Run(":5000")

}

// Define companies slice[dynamic array]
var Companies = []Company{
	{Location: "USA", Name: "Google", CEO: "Sundar Pichai", ID: "1"},
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to tech company API"})
}

func GetCompanies(c *gin.Context) {
	if len(Companies) < 1 {
		c.JSON(http.StatusOK, gin.H{"message": "No companies data"})
		return
	}
	c.JSON(http.StatusOK, Companies)
}

func GetCompany(c *gin.Context) {
	requestID := c.Param("id")

	for _, company := range Companies {
		if company.ID == requestID {
			c.JSON(http.StatusOK, company)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"statusCode": 404, "message": "Company not found"})
}

func PostCompany(c *gin.Context) {
	var newCompany Company
	if err := c.BindJSON(&newCompany); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	Companies = append(Companies, newCompany)
	c.JSON(http.StatusCreated, newCompany)
}

func EditCompany(c *gin.Context) {
	var newCompany Company
	if err := c.BindJSON(&newCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	for index, company := range Companies {
		if company.ID == newCompany.ID {
			company.Location = newCompany.Location
			company.Name = newCompany.Name
			company.CEO = newCompany.CEO
			// Insert into company slice
			Companies = append(Companies[:index], Companies[index+1:]...)
			Companies = append(Companies, newCompany)
			c.JSON(http.StatusOK, newCompany)
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "No company found"})
		}

	}
}

func DeleteCompany(c *gin.Context) {
	requestID := c.Param("id")
	for index, company := range Companies {

		if company.ID == requestID {

			Companies = append(Companies[:index], Companies[index+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Company has been deleted",
			})
			return
		}
	}

	c.JSON(http.StatusOK, Companies)
}
