package main

import (
	"context"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	ginConfig "git.ai-space.tech/coursework/backend/internal/presentation/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	ginApp := gin.Default()
	ctx := context.Background()

	connections := infrastructure.NewConnections(ctx)
	defer connections.Close()

	ginApp.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	ginConfig.RegisterRoutes(ginApp, connections)

	err := ginApp.Run("localhost:8080")
	if err != nil {
		log.Fatalf("Failed to start application: %v\n", err)
	}
}
