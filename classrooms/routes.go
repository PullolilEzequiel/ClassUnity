package classroom

import (
	"github.com/gin-gonic/gin"
)

func ImportClassRoomRoute(router *gin.Engine) {
	// Grupo de rutas para operaciones de usuario
	r := router.Group("/classrooms")

	/// FindAllClassrooms
	r.GET("", func(c *gin.Context) {
		c.String(200, "ok")
	})

}
