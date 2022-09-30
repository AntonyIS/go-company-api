package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	h "github.com/AntonyIS/GO-REST-API-1/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHomeRoute(t *testing.T) {
	mockResponse := `{"message":"Welcome to tech company API"}`
	r := SetUpRouter()
	r.GET("/", h.Home)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCompanies(t *testing.T) {
	mockResponse := []h.Company{
		{Location: "USA", Name: "Google", CEO: "Sundar Pichai", ID: "1"},
	}
	r := SetUpRouter()
	r.GET("/companies", h.GetCompanies)
	req, _ := http.NewRequest("GET", "/companies", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var companies []h.Company
	json.Unmarshal(w.Body.Bytes(), &companies)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse, companies)
}

func TestGetCompany(t *testing.T) {
	mockResponse := h.Company{Location: "USA", Name: "Google", CEO: "Sundar Pichai", ID: "1"}
	r := SetUpRouter()
	r.GET("/companies/:id", h.GetCompany)

	req, _ := http.NewRequest("GET", "/companies/1", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var company h.Company

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
	company := h.Company{Location: "USA", Name: "Google", CEO: "Sundar Pichai", ID: "1"}
	r := SetUpRouter()
	r.POST("/companies", h.PostCompany)

	js, _ := json.Marshal(company)
	req, _ := http.NewRequest("POST", "/companies", bytes.NewBuffer(js))

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestEditCompany(t *testing.T) {
	company := h.Company{Location: "USA", Name: "Google", CEO: "Sundar Pichai", ID: "1"}
	r := SetUpRouter()
	r.PUT("/companies/:id", h.EditCompany)
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
	r.DELETE("/companies/:id", h.DeleteCompany)

	req, _ := http.NewRequest("DELETE", "/companies/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	var companies []h.Company
	json.Unmarshal(w.Body.Bytes(), &companies)

	assert.Equal(t, len(companies), 0)
	assert.Equal(t, http.StatusOK, w.Code)

}
