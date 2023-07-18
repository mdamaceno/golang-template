package web

import (
    "github.com/gin-gonic/gin"

    "app/web/controller"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/", controller.HomeController["index"])

    userRoutes := r.Group("/users")
    userRoutes.GET("/", controller.UsersController["index"])
    userRoutes.GET("/:id", controller.UsersController["show"])
    userRoutes.GET("/:id/edit", controller.UsersController["edit"])
    userRoutes.GET("/new", controller.UsersController["new"])
    userRoutes.POST("/", controller.UsersController["create"])
    userRoutes.PATCH("/:id", controller.UsersController["update"])
    userRoutes.DELETE("/:id", controller.UsersController["destroy"])
}
