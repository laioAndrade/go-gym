package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laioandrade/go-gym/models"
	"github.com/laioandrade/go-gym/repositories"
)

type CreateAlunoInput struct {
	Name string `json:"nome" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"senha" binding:"required"`
	Age  uint8	`json:"idade" binding:"required"`
	Gender string `json:"sexo" binding:"required"`
}

type UpdateAlunoInput struct {
	Name string `json:"nome"`
	UserName string `json:"username"`
	Password string `json:"senha"`
	Age  uint8	`json:"idade"`
	Gender string `json:"sexo"`
}

func CreateAluno(c *gin.Context){
	var input CreateAlunoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aluno := models.Aluno{
	Name: input.Name,
	UserName: input.UserName,
	Password: input.Password,
	Age: input.Age,
	Gender: input.Gender,

	}

	if err := repositories.CreateAluno(&aluno); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "usuario cadastrado com sucesso"})
}

func GetAlunos(c *gin.Context) {
	alunos := repositories.GetAlunos()

	c.JSON(http.StatusOK, gin.H{"data": alunos})
}

func GetAluno(c *gin.Context) {
	var aluno models.Aluno

	if repositories.GetAluno(&aluno, c.Param("id")) != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": aluno})
}

func UpdateAluno(c *gin.Context) {
	var aluno models.Aluno

	if repositories.GetAluno(&aluno, c.Param("id")) != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "error"})
		return
	}

	var input models.Aluno
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	
	if repositories.UpdateAluno(&aluno, input) != nil {
		c.JSON(http.StatusConflict, gin.H{"error" : "Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": aluno})
}

func DeleteAluno(c *gin.Context) {
	var aluno models.Aluno

	if repositories.GetAluno(&aluno, c.Param("id")) != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "error"})
		return
	}

	if repositories.DeleteAluno(&aluno)!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}