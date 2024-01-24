package main

import (
	classroom "class_unity/classrooms"
	"class_unity/users"

	"github.com/gin-gonic/gin"
)

func main() {
	a := gin.Default()
	users.ImportUserRoutes(a)
	classroom.ImportClassRoomRoute(a)
	a.Run()
}
