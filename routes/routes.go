package routes

import (
	"github.com/Renan-Parise/codium-mail/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	emailController := controllers.NewEmailController()

	mailRoutes := router.Group("/mail")
	{
		mailRoutes.POST("/send", emailController.Send)
	}

	return router
}
