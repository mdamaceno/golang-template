package web

import (
    "github.com/gin-gonic/gin"

    "app/web/controller"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/", controller.HomeController(r, "index"))
    r.GET("/users", controller.UsersController(r, "index"))
}
