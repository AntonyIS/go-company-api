package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHomeRoute(t *testing.T) {
	mockResponse := `{"message":"Welcome to tech company API"}`
	r := SetUpRouter()
	r.GET("/", Home)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCompanies(t *testing.T) {
	mockResponse := []Company{
		{Location: "USA", Name: "Google", CEO: "Sundar Pichai", ID: "1"},
	}
	r := SetUpRouter()
	r.GET("/companies", GetCompanies)
	req, _ := http.NewRequest("GET", "/companies", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var companies []Company
	json.Unmarshal(w.Body.Bytes(), &companies)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse, companies)
}

func TestGetCompany(t *testing.T) {
	mockResponse := Company{Location: "USA", Name: "Google", CEO: "Sundar Pichai", ID: "1"}
	r := SetUpRouter()
	r.GET("/companies/:id", GetCompany)

	req, _ := http.NewRequest("GET", "/companies/1", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var company Company

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &company); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, mockResponse, company)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestPostCompany(t *testing.T) {
	company := Company{Location: "USA", Name: "Google", CEO: "Sundar Pichai", ID: "1"}
	r := SetUpRouter()
	r.POST("/companies", PostCompany)

	js, _ := json.Marshal(company)
	req, _ := http.NewRequest("POST", "/companies", bytes.NewBuffer(js))

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestEditCompany(t *testing.T) {
	company := Company{Location: "USA", Name: "Google", CEO: "Sundar Pichai", ID: "1"}
	r := SetUpRouter()
	r.PUT("/companies/:id", EditCompany)
	js, _ := json.Marshal(company)
	reqFound, _ := http.NewRequest("PUT", "/companies/"+company.ID, bytes.NewBuffer(js))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, reqFound)
	assert.Equal(t, http.StatusOK, res.Code)

	reqNotFound, _ := http.NewRequest("PUT", "/companies/50", bytes.NewBuffer(js))
	res = httptest.NewRecorder()
	r.ServeHTTP(res, reqNotFound)
	assert.Equal(t, http.StatusNotFound, 404)
}

func TestDeleteCompany(t *testing.T) {

	r := SetUpRouter()
	r.DELETE("/companies/:id", DeleteCompany)

	req, _ := http.NewRequest("DELETE", "/companies/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	var companies []Company
	json.Unmarshal(w.Body.Bytes(), &companies)

	assert.Equal(t, len(companies), 0)
	assert.Equal(t, http.StatusOK, w.Code)

}
