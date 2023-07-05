package web

import (
    "github.com/gin-gonic/gin"

    "app/web/controller"
)

func SetupRoutes(r *gin.Engine) *gin.Engine {
    r.GET("/", controller.HomeIndex(r))

    return r
}
