package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/wecercle/test-devjr-fullstack-cercle/cmd/api/docs"
	"github.com/wecercle/test-devjr-fullstack-cercle/cmd/api/v1/app"
)

// setupCORS é uma função auxiliar para configurar o middleware CORS no router Gin, permitindo solicitações de qualquer origem e métodos comuns.
func setupCORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization"},
	}))
}

// start é a função principal que inicializa o router Gin, configura CORS, Swagger e as rotas da aplicação, e inicia o servidor na porta 8080.
func start() {
	router := gin.Default()

	setupCORS(router)
	setupSwagger(router)
	app.Routes{}.SetupRouterV1(router)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func main() {
	start()
}

// setupSwagger é uma função auxiliar para configurar as rotas do Swagger UI no router Gin, permitindo acesso à documentação da API em /swagger/*any.
func setupSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))
}
