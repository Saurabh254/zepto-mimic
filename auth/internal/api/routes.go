package api

import (
	"github.com/gin-gonic/gin"
	"github.com/saurabh254/zepto-mimic/auth/internal/api/handlers"
)

func RegisterRoutes(rg *gin.RouterGroup) {

	handlers.RegisterAuthRoutes(rg.Group("/auth"))
	handlers.RegisterUserRoutes(rg.Group("/user"))

}
