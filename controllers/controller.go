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

func BuscaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	databases.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func DeletarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	databases.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Aluno deletado com sucesso",
	})
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	databases.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}
	databases.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	databases.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
