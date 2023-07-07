package bootstrap

import (
    "github.com/gin-gonic/gin"

    "app/web"
    "app/web/view"
)

func Serve() {
    r := gin.New()
    r.HTMLRender = view.SetViews()

    setupAssets(r)
    setupDefaultMiddleware(r)

    web.SetupRoutes(r)

    r.Run()
}

func setupAssets(r *gin.Engine) *gin.Engine {
    r.Static("/public", "./public")

    return r
}

func setupDefaultMiddleware(r *gin.Engine) *gin.Engine {
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    return r
}
