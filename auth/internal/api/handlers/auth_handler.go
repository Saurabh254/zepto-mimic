package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(c *gin.RouterGroup) {
	c.POST("/login", Login)
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Authenticate user and return JWT token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Failure		400	{object}	map[string]interface{}
//	@Router			/api/v1/auth/login [post]
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "login",
	})
}
