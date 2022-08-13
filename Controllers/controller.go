package Controllers

import "github.com/gin-gonic/gin"

func ExibeTodosOsAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Murilo",
	})
}

func RecuperarAluno(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}
