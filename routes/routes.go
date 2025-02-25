package routes

import (
	"customer/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine){
	routes := router.Group("/polling")
	{
		routes.GET("/users",controllers.CheckListUsers)
		// routes.GET("",)
		// routes.GET("",)
		// routes.GET("",)
	}
}