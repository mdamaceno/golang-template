package main

import (
    "path/filepath"
    "github.com/gin-gonic/gin"

    "app/config"
)

func main() {
    r := gin.New()
    r.LoadHTMLGlob(filepath.Join("web", "resources", "views", "**", "*.html"))

    r = config.SetupDefaultMiddleware(r)
    r = config.SetupRoutes(r)

    r.Run()
}
