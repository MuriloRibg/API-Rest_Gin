package Routes

import (
	"github.com/MuriloRibg/API-Rest_Gin.git/Controllers"
	"github.com/MuriloRibg/API-Rest_Gin.git/Middleware"
	"github.com/MuriloRibg/API-Rest_Gin.git/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	_ "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func HandleRequests() {
	r := gin.Default()

	//swag init --pd --parseInternal --parseDepth 1
	//Informações do swagger
	docs.SwaggerInfo.Title = "Api do Murilove"
	docs.SwaggerInfo.Description = "Api REST com Gin e Gorm"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "18.230.75.0:8080/v1"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.Use(Middleware.CORSMiddleware())
	r.GET("/alunos", Controllers.Listar)
	r.GET("/aluno/:id", Controllers.Recuperar)
	r.GET("/aluno/cpf/:cpf", Controllers.RecuperPorCPF)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/aluno", Controllers.Inserir)
	r.DELETE("/aluno/:id", Controllers.Deletar)
	r.PUT("/aluno/:id", Controllers.Editar)

	err := r.Run()
	if err != nil {
		panic(err.Error())
	}
}
