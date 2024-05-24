package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/models"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

func CriaNovoAluno(c *gin.context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {	
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()
		})
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}