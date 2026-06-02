package main

import (
	_ "github.com/saurabh254/zepto-mimic/auth/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/saurabh254/zepto-mimic/auth/internal/api"
	"github.com/saurabh254/zepto-mimic/auth/internal/core/config"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Auth API
// @version 1.0
// @description Auth service
// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.LoadConfig()
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	r.Use(cors.Default())

	r.StaticFile("/swagger.yaml", "./internal/openapi/api.yaml")

	api.RegisterRoutes(r.Group("/api/v1"))
	api.RegisterHealthRoutes(r.Group("/healthz"))

	r.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run(cfg.Address())
}
