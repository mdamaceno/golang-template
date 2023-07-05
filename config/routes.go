package config

import (
    "github.com/gin-gonic/gin"

    "app/web/controller"
)

func SetupRoutes(r *gin.Engine) *gin.Engine {
    r.GET("/", controller.HomeIndex)

    return r
}
