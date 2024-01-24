package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ImportUserRoutes(router *gin.Engine) {
	// Grupo de rutas para operaciones de usuario
	r := router.Group("/users")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Operaciones de usuario",
		})
	})

}
