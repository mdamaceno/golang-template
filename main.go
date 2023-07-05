package main

import (
    "github.com/gin-gonic/gin"

    "app/config"
)

func main() {
    r := gin.New()

    r = config.SetupAssets(r)
    r = config.SetupLayouts(r)
    r = config.SetupDefaultMiddleware(r)
    r = config.SetupRoutes(r)

    r.Run()
}
