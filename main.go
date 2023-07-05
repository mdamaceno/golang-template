package main

import (
    "path/filepath"
    "github.com/gin-gonic/gin"

    "app/config"
)

func main() {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Static("/public", "./public")
    r.LoadHTMLGlob(filepath.Join("web", "resources", "views", "**", "*.html"))

    r = config.SetupRoutes(r)

    r.Run()
}
