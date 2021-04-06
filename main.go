package main

import (
	"github.com/gin-gonic/gin"
	"github.com/laioandrade/go-gym/controllers"
	"github.com/laioandrade/go-gym/database"
)

func main(){
	r := gin.Default()
	database.Initialize()

	r.POST("/", controllers.CreateAluno)
	r.GET("/", controllers.GetAlunos)
	r.GET("/:id", controllers.GetAluno)
	r.PATCH("/:id", controllers.UpdateAluno)
	r.DELETE("/:id", controllers.DeleteAluno)




	r.Run(":3333")
}