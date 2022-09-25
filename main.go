package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	Ldate     = 1 << iota // the date in the local time zone: 2009/01/23
	YYYYMMDD  = "2006-01-02"
	HHMMSS12h = "3:04:05 PM"
)

type Company struct {
	Location string `json:"location"`
	Name     string `json:"name"`
	CEO      string `json:"ceo"`
	ID       string `json:"id"`
}

// Define companies slice[dynamic array]
var companies = []Company{
	{Location: "USA", Name: "Google", CEO: "Sundar Pichai", ID: "1"},
}

func main() {
	setUpLogger()
	router := gin.Default()

	router.GET("/", Home)
	router.GET("/companies", GetCompanies)
	router.GET("/companies/:id", GetCompany)
	router.POST("/companies", PostCompany)
	router.PUT("/companies/:id", EditCompany)
	router.DELETE("/companies/:id", DeleteCompany)
	// Run server
	router.Run(":5000")
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to tech company API"})
}

func GetCompanies(c *gin.Context) {
	c.JSON(http.StatusOK, companies)
}

func GetCompany(c *gin.Context) {
	requestID := c.Param("id")

	for _, company := range companies {
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
	companies = append(companies, newCompany)
	c.JSON(http.StatusCreated, newCompany)
}

func EditCompany(c *gin.Context) {
	var newCompany Company
	if err := c.BindJSON(&newCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	for index, company := range companies {
		if company.ID == newCompany.ID {
			company.Location = newCompany.Location
			company.Name = newCompany.Name
			company.CEO = newCompany.CEO
			// Insert into company slice
			companies = append(companies[:index], companies[index+1:]...)
			companies = append(companies, newCompany)
			c.JSON(http.StatusOK, newCompany)
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "No company found"})
		}

	}
}

func DeleteCompany(c *gin.Context) {
	requestID := c.Param("id")
	for index, company := range companies {
		if company.ID == requestID {
			companies = append(companies[:index], companies[index+1:]...)
			c.JSON(http.StatusOK, companies)
			return
		}
	}
	c.JSON(http.StatusOK, companies)
}

func setUpLogger() {
	// Set up logger
	flag := log.Ldate
	datetime := time.Now().UTC().Format(YYYYMMDD+" "+HHMMSS12h) + ": "
	log.SetFlags(flag)
	log.SetPrefix(datetime)
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger := log.New(file, "", flag)
		logger.SetPrefix("FATAL: " + datetime)
		logger.Println(err)
	}
	defer file.Close()
	logger := log.New(file, "", flag)
	logger.SetPrefix("INFO : " + datetime)

	mw := io.MultiWriter(os.Stdout, file)

	logger.SetOutput(mw)
	logger.Println("Starting the company API")
}
