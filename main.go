package main

import (
    "github.com/gin-gonic/gin"

    "app/config"
)

func main() {
    r := gin.New()

    config.SetupAssets(r)
    config.SetupLayouts(r)
    config.SetupDefaultMiddleware(r)
    config.SetupRoutes(r)

    r.Run()
}
