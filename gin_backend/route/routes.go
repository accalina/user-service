package route

import (
	"gin_backend/controller"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.GET("/", controller.Home)

	// user
	r.GET("/user", controller.ListUser)
	r.GET("/user/:id", controller.RetriveUser)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:id", controller.UpdateUser)
	r.DELETE("/user/:id", controller.DeleteUser)

	return r
}
