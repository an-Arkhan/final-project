package router

import (
	"final-project/app"
	"final-project/controller"
	"final-project/middleware"
	"final-project/repository"
	"final-project/service"

	"github.com/gin-gonic/gin"
)

func SocialMediaRouter(router *gin.Engine) {
	db := app.NewDB()

	repo := repository.NewSocialMediaRepository(db)
	srv := service.NewSocialMediaService(repo)
	ctrl := controller.NewSocialMediaController(srv)

	socialMedia := router.Group("/socialmedias", middleware.AuthMiddleware())

	{
		socialMedia.GET("/", ctrl.GetAll)
		socialMedia.POST("/", ctrl.Create)
		{
			socialMedia.PUT("/:id", middleware.SocialMediaMiddleware(srv), ctrl.Update)
			socialMedia.DELETE("/:id", middleware.SocialMediaMiddleware(srv), ctrl.Delete)
		}
	}
}
