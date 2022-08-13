package Routes

import (
	"github.com/MuriloRibg/API-Rest_Gin.git/Controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", Controllers.ExibeTodosOsAlunos)
	r.GET("/:nome", Controllers.RecuperarAluno)
	r.Run()
}
