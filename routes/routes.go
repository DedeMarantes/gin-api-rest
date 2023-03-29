package routes

import (
	"github.com/DedeMarantes/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	app := gin.Default()
	app.GET("/alunos", controllers.ExibeAlunos)
	app.GET("/:nome", controllers.Saudacao)
	app.GET("/alunos/:id", controllers.BuscaAluno)
	app.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	app.POST("/alunos", controllers.CriaAluno)
	app.PATCH("/alunos/:id", controllers.EditaAluno)
	app.DELETE("/alunos/:id", controllers.DeletarAluno)
	app.Run(":8600")
}
