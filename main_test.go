package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/douglira/alura-golang-gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	return r
}

func TestHealthcheckWithParameterAndBody(t *testing.T) {
	r := SetupRoutes()
	r.GET("/api/healthcheck/:ping", controllers.Healthcheck)
	req, _ := http.NewRequest(http.MethodGet, "/api/healthcheck/ping", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "Should be equal")
	mockResponse := `{"data":"pong"}`
	resBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, mockResponse, string(resBody))
}
