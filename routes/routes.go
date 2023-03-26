package routes

import (
	"github.com/DedeMarantes/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	app := gin.Default()
	app.GET("/alunos", controllers.ExibeAlunos)
	app.GET("/:nome", controllers.Saudacao)
	app.POST("/alunos", controllers.CriaAluno)
	app.Run(":8600")
}
