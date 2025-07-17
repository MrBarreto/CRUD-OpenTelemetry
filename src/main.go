// src/main.go
package main

import (
	"log"

	"github.com/MrBarreto/CRUD-OpenTelemetry/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	router.Run(":8080")
}
