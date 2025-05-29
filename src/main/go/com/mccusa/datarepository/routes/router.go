package routes

import (
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	// Public
	//r.POST("/register", controller.Register)
	//r.POST("/login", controller.Login)

	// Protected
	//auth := r.Group("/api", middleware.JWTAuth(config.Cfg.JWTSecret))
	{
		//auth.GET("/users", controllers.GetUsers)
		//auth.GET("/users/:id", controllers.GetUser)
		// add POST/PUT/DELETE here…
	}

	return r
}
