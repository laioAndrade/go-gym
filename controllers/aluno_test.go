package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/laioandrade/go-gym/database"
)

func TestGetAluno(t *testing.T) {
	router := gin.Default()
	database.Initialize()

	router.GET("/", GetAlunos)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("requisição retornou status errado: obtido %v  esperado %v", status, http.StatusOK)
	}

	expected := `{"data":[{"ID":0,"DeletedAt":null,"id":1,"nome":"laio","username":"laioAlencar","senha":"123456","idade":23,"sexo":"masculino","CreatedAt":"2021-04-06T19:05:48.08-03:00","UpdatedAt":"2021-04-06T19:31:04.303-03:00"}]}`

	if w.Body.String() != expected {
		t.Errorf("requisição retornou body inesperado: : obtido %v  esperado %v", w.Body.String(), expected)
	}
}