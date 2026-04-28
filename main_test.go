package main

import (
	"go-rest-gin/controllers"
	"go-rest-gin/database"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRoutesTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()

	return rotas
}

func TestVerifyStatusCodeStudents(t *testing.T) {
	database.Connection()
	r := SetupRoutesTest()
	r.GET("/students/:id", controllers.GetById)

	req, _ := http.NewRequest("GET", "/students/1", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	mockResponse := `{
		"ID": 1,
		"CreatedAt": "2026-04-27T23:43:28.390987-03:00",
		"UpdatedAt": "2026-04-27T23:43:28.390987-03:00",
		"DeletedAt": null,
		"nome": "Bruno Moura",
		"cpf": "000",
		"rg": "000"
	}`

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		assert.Error(t, err, "Não foi possível ler o body da request")
	}

	assert.JSONEq(t, mockResponse, string(responseBody))
}
