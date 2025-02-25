package main

import (
	"customer/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	PORT := ":8081"

	routes.UserRouter(r)

	r.Run(PORT)

}