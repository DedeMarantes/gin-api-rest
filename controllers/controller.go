package controllers

import (
	"fmt"
	"net/http"

	"github.com/DedeMarantes/gin-api-rest/databases"
	"github.com/DedeMarantes/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func ExibeAlunos(c *gin.Context) {
	var alunos []models.Aluno
	databases.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"mensagem": fmt.Sprintf("E ai %s, tudo beleza?", nome),
	})
}

func CriaAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databases.DB.Create(&aluno)
	c.JSON(http.StatusCreated, aluno)

}
