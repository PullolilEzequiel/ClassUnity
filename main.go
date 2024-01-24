package main

import (
	"class_unity/users"

	"github.com/gin-gonic/gin"
)

func main() {
	a := gin.Default()
	users.ImportUserRoutes(a)
	a.Run()
}
