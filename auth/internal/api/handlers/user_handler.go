package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(c *gin.RouterGroup) {
	c.GET("/profile", UserProfile)
}

func UserProfile(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"msg": "user profile",
	})

}
