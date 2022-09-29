# Go REST API
A simple Golang REST API that exposes CRUD endpoints. These endpoints expose technology company data to the client.
This REST API makes use of the Go Gin framework to handle requests from the client. It also has has unit test for the exposed endpoints.
Docker has also been used to containerize the application. The image of the API will be uploaded soon to Docker Hub.
As much as I wanted this to be a minimal application, below are the features that I would like to impelement.
* Hexagonal Architecure
* Databases
* Caching with Redis
* Logging


## Installation.
* Clone this application in your working directory
    * * $git clone https://github.com/AntonyIS/Go-REST-API-1
* Build a docker images into your machine. Add the below command in you terminal.
    * * $docker build -t test-api .
* Run the application. This app exposes port 5000. 
    * * $ docker run -it -p 5000:5000 test-api
* Access the API endpoints
    * * Using Postman or browser, hit this endpoint http://127.0.0.1:5000/companies

