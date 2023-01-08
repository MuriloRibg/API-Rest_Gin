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
	r.LoadHTMLGlob("Templates/*")
	r.Static("/Assets", "./Assets")

	//swag init --pd --parseInternal --parseDepth 1
	//Informações do swagger
	docs.SwaggerInfo.Title = "Api do Murilove"
	docs.SwaggerInfo.Description = "Api REST com Gin e Gorm"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.Use(Middleware.CORSMiddleware())
	r.GET("/alunos", Controllers.Listar)
	r.GET("/aluno/:id", Controllers.Recuperar)
	r.GET("/aluno/cpf/:cpf", Controllers.RecuperPorCPF)
	r.POST("/aluno", Controllers.Inserir)
	r.DELETE("/aluno/:id", Controllers.Deletar)
	r.PUT("/aluno/:id", Controllers.Editar)

	r.GET("/inicio", Controllers.ExibePaginaInicial)
	r.GET("/inicio/alunos", Controllers.ExibePaginaAlunos)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(Controllers.RotaNaoEncontrada)
	err := r.Run()
	if err != nil {
		panic(err.Error())
	}
}
