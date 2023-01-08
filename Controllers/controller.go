package Controllers

import (
	"fmt"
	"github.com/MuriloRibg/API-Rest_Gin.git/Database"
	"github.com/MuriloRibg/API-Rest_Gin.git/Models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"net/http"
)

// Listar 		 godoc
// @Summary      Mostrar todos os alunos
// @Description  Rota para buscar todos os alunos
// @Tags         Aluno
// @Accept       json
// @Produce      json
// @Success      200  {object}  Models.Aluno
// @Router       /alunos [get]
func Listar(c *gin.Context) {
	var Alunos []Models.Aluno
	Database.DB.Find(&Alunos)
	c.JSON(http.StatusOK, Alunos)
}

// Recuperar 	 godoc
// @Summary      Recupar um Aluno
// @Description  get aluno por ID
// @Tags         Aluno
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Aluno ID"
// @Success      200  {array}   Models.Aluno
// @Failure      404  {object}  httputil.HTTPError
// @Router       /aluno/{id} [get]
func Recuperar(c *gin.Context) {
	id := c.Params.ByName("id")

	var aluno Models.Aluno
	Database.DB.Find(&aluno, id)
	if aluno.ID > 0 {
		c.JSON(http.StatusOK, aluno)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"Error": fmt.Sprintf("Aluno com id = %s, não encontrado!", id),
	})
}

// Inserir   	 godoc
// @Summary      Inserir um novo aluno
// @Description  post aluno
// @Tags         Aluno
// @Accept       json
// @Produce      json
// @Param        aluno	body	Models.Aluno  true  "Modelo do aluno"
// @Success      200  {array}   Models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Router       /aluno [post]
func Inserir(c *gin.Context) {
	var aluno Models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := Models.ValidarDados(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	Database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// Deletar 		 godoc
// @Summary      Deletar um aluno
// @Description  delete aluno por ID
// @Tags         Aluno
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Aluno ID"
// @Success      200
// @Failure      400  {object}  httputil.HTTPError
// @Router       /aluno/{id} [delete]
func Deletar(c *gin.Context) {
	id := c.Params.ByName("id")
	var Aluno Models.Aluno
	resp := Database.DB.Delete(&Aluno, id)
	if resp.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"Mesagem": fmt.Sprintf("Aluno com o id = %s, apagado com sucesso!", id),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"Error": resp.Error.Error(),
	})
}

// Editar 		 godoc
// @Summary      Editar um aluno
// @Description  put aluno por id
// @Tags         Aluno
// @Accept       json
// @Produce      json
// @Param        aluno	body	Models.Aluno	true	"Modelo de aluno"
// @Success      200  {array}   Models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Router       /aluno/{id} [put]
func Editar(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno Models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := Models.ValidarDados(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	Database.DB.Model(&aluno).Where("id = ?", id).Updates(aluno)
	c.JSON(http.StatusOK, aluno)
}

// RecuperPorCPF godoc
// @Summary      Recupera aluno por CPF
// @Description  get aluno por cpf
// @Tags         Aluno
// @Accept       json
// @Produce      json
// @Param        cpf	path	string	true  "Aluno CPF"
// @Success      200  {array}   Models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Router       /aluno/cpf/{cpf} [get]
func RecuperPorCPF(c *gin.Context) {
	cpf := c.Params.ByName("cpf")
	var aluno Models.Aluno
	Database.DB.Where(&Models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID > 0 {
		c.JSON(http.StatusOK, aluno)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"Error": fmt.Sprintf("Aluno com cpf = %s, não encontrado!", cpf),
	})
}

func ExibePaginaInicial(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"mensagem": "Boas vindas",
	})
}

func ExibePaginaAlunos(c *gin.Context) {
	var alunos []Models.Aluno
	Database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "alunos.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
