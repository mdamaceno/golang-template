package config

import (
    "path/filepath"
    "github.com/gin-gonic/gin"

    "app/web/controller"
)

func SetupRoutes() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Static("/public", "./public")
    r.LoadHTMLGlob(filepath.Join("web", "resources", "views", "**", "*.html"))

    r.GET("/", controller.HomeIndex)

    return r
}
