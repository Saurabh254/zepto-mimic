package api

import "github.com/gin-gonic/gin"

func RegisterHealthRoutes(rg *gin.RouterGroup) {

	rg.GET("/", func(c *gin.Context) {
  		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}
